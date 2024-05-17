package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cli/browser"
	"github.com/cli/oauth"
)

const (
	oauth20Token            = `https://login.live.com/oauth20_token.srf`
	authenticateURL         = `https://user.auth.xboxlive.com/user/authenticate`
	authenticatewithXSTSURL = `https://xsts.auth.xboxlive.com/xsts/authorize`
	loginWithXboxURL        = `https://api.minecraftservices.com/authentication/login_with_xbox`
	getTheprofileURL        = `https://api.minecraftservices.com/minecraft/profile`
)

func MsLoginRefresh(t *MsToken) (*Profile, error) {
	has := false
	if t != nil {
		err := t.Refresh()
		if err == nil {
			has = true
		}
	}
	if !has {
		var err error
		t, err = getToken()
		if err != nil {
			return nil, fmt.Errorf("MsLogin: %w", err)
		}
	}
	xbltoken, uhs, err := getXbltoken(t.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("MsLogin: %w", err)
	}
	xststoken, err := getXSTStoken(xbltoken)
	if err != nil {
		return nil, fmt.Errorf("MsLogin: %w", err)
	}
	AccessToken, err := loginWithXbox(uhs, xststoken)
	if err != nil {
		return nil, fmt.Errorf("MsLogin: %w", err)
	}
	p, err := GetProfile(AccessToken)
	if err != nil {
		return nil, fmt.Errorf("MsLogin: %w", err)
	}
	p.AccessToken = AccessToken
	p.MsToken = *t
	return p, nil
}

func getToken() (*MsToken, error) {
	var t msToken
	f := oauth.Flow{
		Scopes:   []string{"XboxLive.signin", "offline_access"},
		ClientID: "a48a9fad-1702-46d7-8ee9-42b857ad292d",
		Host: &oauth.Host{
			DeviceCodeURL: "https://login.microsoftonline.com/consumers/oauth2/v2.0/devicecode",
			TokenURL:      "https://login.microsoftonline.com/consumers/oauth2/v2.0/token",
		},
		BrowseURL: func(s string) error {
			fmt.Println("If not opening a web browser, try to open " + s)
			err := browser.OpenURL(s)
			if err != nil {
				fmt.Println("open browser failed: " + err.Error())
			}
			return nil
		},
		HTTPClient: c,
	}
	token, err := f.DeviceFlow()
	if err != nil {
		return nil, fmt.Errorf("getToken: %w", err)
	}
	t.AccessToken = token.Token
	t.RefreshToken = token.RefreshToken
	m := MsToken{}
	m.parse(t)
	return &m, nil
}

func getXbltoken(token string) (xblToken, uhs string, err error) {
	msg := `{"Properties": {"AuthMethod": "RPS","SiteName": "user.auth.xboxlive.com","RpsTicket": "d=` + jsonEscape(token) + `"},"RelyingParty": "http://auth.xboxlive.com","TokenType": "JWT"}`
	b, err := httPost(authenticateURL, msg, `application/json`)
	if err != nil {
		return "", "", fmt.Errorf("getXbltoken: %w", err)
	}
	m := msauth{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", "", fmt.Errorf("getXbltoken: %w", err)
	}
	if len(m.DisplayClaims.Xui) < 1 {
		return "", "", ErrToken
	}
	return m.Token, m.DisplayClaims.Xui[0].Uhs, nil
}

func getXSTStoken(xblToken string) (string, error) {
	msg := `{
		"Properties": {
			"SandboxId": "RETAIL",
			"UserTokens": [
				"` + jsonEscape(xblToken) + `" 
			]
		},
		"RelyingParty": "rp://api.minecraftservices.com/",
		"TokenType": "JWT"
	 }`
	b, err := httPost(authenticatewithXSTSURL, msg, `application/json`)
	if err != nil {
		e := ErrHTTPCode{}
		if errors.As(err, &e) && e.code == 401 {
			m := map[string]interface{}{}
			err1 := json.Unmarshal([]byte(e.msg), &m)
			if err1 != nil {
				return "", fmt.Errorf("getXSTStoken: %w", err)
			}
			code, ok := m["XErr"]
			if ok {
				c, _ := code.(float64)
				switch int(c) {
				case 2148916238:
					return "", ErrChild
				case 2148916233:
					return "", ErrXboxNotLinked
				}
			}
		}
		return "", fmt.Errorf("getXSTStoken: %w", err)
	}
	m := msauth{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", fmt.Errorf("getXSTStoken: %w", err)
	}
	return m.Token, nil
}

var (
	ErrChild         = errors.New("child account or not added to a Family")
	ErrXboxNotLinked = errors.New("xbox account not exist")
)

func loginWithXbox(uhs string, xstsToken string) (string, error) {
	msg := `{"identityToken": "XBL3.0 x=` + jsonEscape(uhs) + `;` + jsonEscape(xstsToken) + `"}`
	b, err := httPost(loginWithXboxURL, msg, `application/json`)
	if err != nil {
		return "", fmt.Errorf("loginWithXbox: %w", err)
	}
	t := msToken{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return "", fmt.Errorf("loginWithXbox: %w", err)
	}
	return t.AccessToken, nil
}

func GetProfile(authorization string) (*Profile, error) {
	reqs, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, getTheprofileURL, nil)
	if err != nil {
		return nil, fmt.Errorf("getProfile: %w", err)
	}
	reqs.Header.Set("Authorization", "Bearer "+authorization)
	rep, err := c.Do(reqs)
	if rep != nil {
		defer func() {
			if err := rep.Body.Close(); err != nil {
				panic(err)
			}
		}()
	}
	if err != nil {
		return nil, fmt.Errorf("getProfile: %w", err)
	}
	b, err := io.ReadAll(rep.Body)
	if err != nil {
		return nil, fmt.Errorf("getProfile: %w", err)
	}
	p := Profile{
		AccessToken: authorization,
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, fmt.Errorf("getProfile: %w", err)
	}
	if p.ID == "" {
		return nil, ErrProfile
	}
	return &p, nil
}

type Profile struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AccessToken string
	MsToken     MsToken
}

type msauth struct {
	DisplayClaims displayClaims `json:"DisplayClaims"`
	IssueInstant  string        `json:"IssueInstant"`
	NotAfter      string        `json:"NotAfter"`
	Token         string        `json:"Token"`
}

type displayClaims struct {
	Xui []xui `json:"xui"`
}

type xui struct {
	Uhs string `json:"uhs"`
}

var (
	ErrCode    = errors.New("code invalid")
	ErrToken   = errors.New("token invalid")
	ErrProfile = errors.New("do not have game")
)

func httPost(url, msg, contentType string) ([]byte, error) {
	reqs, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, url, strings.NewReader(msg))
	if err != nil {
		return nil, fmt.Errorf("httPost: %w", err)
	}
	reqs.Header.Set("Content-Type", contentType)
	reqs.Header.Set("Accept", "*/*")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	rep, err := c.Do(reqs)
	if rep != nil {
		defer func() {
			if err := rep.Body.Close(); err != nil {
				panic(err)
			}
		}()
	}
	if err != nil {
		return nil, fmt.Errorf("httPost: %w", err)
	}
	b, err := io.ReadAll(rep.Body)
	if err != nil {
		return nil, fmt.Errorf("httPost: %w", err)
	}
	if rep.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httPost: %w", ErrHTTPCode{
			code: rep.StatusCode,
			msg:  string(b),
		})
	}
	return b, nil
}

var c = &http.Client{
	Timeout:   15 * time.Second,
	Transport: Transport,
}

func jsonEscape(s string) string {
	b, err := json.Marshal(&s)
	if err != nil {
		panic(err)
	}
	r := []rune(string(b))
	if len(r) == 0 {
		return ""
	}
	if r[0] == '"' {
		r = r[1:]
	}
	if r[len(r)-1] == '"' {
		r = r[:len(r)-1]
	}
	return string(r)
}
