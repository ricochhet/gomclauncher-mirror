package auth

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var Transport = http.DefaultTransport.(*http.Transport).Clone()

func init() {
	Transport.TLSClientConfig = &tls.Config{
		Renegotiation: tls.RenegotiateOnceAsClient,
	}
	// microsoft auth neet this
}

//lint:ignore ST1008 // ...
//nolint:stylecheck // ...
func post(apiAddress, endpoint string, payload []byte) ([]byte, error, int) {
	var api string
	if apiAddress != "https://sessionserver.mojang.com" {
		var err error
		api, err = url.JoinPath(apiAddress, "/authserver")
		if err != nil {
			return nil, fmt.Errorf("post: %w", err), 0
		}
	}
	if api == "" {
		api = "https://sessionserver.mojang.com"
	}
	h, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, api+"/"+endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Accept", "*/*")
	h.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	c := &http.Client{
		Timeout:   5 * time.Second,
		Transport: Transport,
	}
	rep, err := c.Do(h)
	if rep != nil {
		defer func() {
			if err := rep.Body.Close(); err != nil {
				panic(err)
			}
		}()
	}
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	b, err := io.ReadAll(rep.Body)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	return b, nil, rep.StatusCode
}
