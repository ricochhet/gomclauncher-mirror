package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/xmdhs/gomclauncher/auth"
	aflag "github.com/xmdhs/gomclauncher/flag"
	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func main() {
	if v {
		version()
	}
	if tidy {
		f.Tidy()
		return
	}
	if f.Proxy != "" {
		proxy, err := url.Parse(f.Proxy)
		if err != nil {
			panic(err)
		}
		auth.Transport.Proxy = http.ProxyURL(proxy)
	}
	if credit {
		credits()
	}
	if f.Verlist != "" {
		fmt.Println(strings.Join(f.Arunlist(), "\n"))
	}
	if f.Verlistfabric != "" {
		v, err := f.Arunfabriclist()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Join(v, "\n"))
	}
	if f.Verlistquilt != "" {
		v, err := f.Arunquiltlist()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Join(v, "\n"))
	}
	if f.Verlistpaper != "" {
		v, err := f.Arunpaperlist()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Join(v, "\n"))
	}
	if f.Verlistjava {
		v, err := f.Arunjavalist()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Join(v, "\n"))
	}
	if f.Runlist {
		s := internal.Find(launcher.Minecraft + `/versions`)
		for _, v := range s {
			if aflag.Test(launcher.Minecraft + `/versions/` + v + `/` + v + ".json") {
				fmt.Println(v)
			}
		}
	}
	if f.Runlistpaper {
		s := internal.Find(launcher.Minecraft + `/servers`)
		for _, v := range s {
			if aflag.Testservers(launcher.Minecraft + `/servers/` + v + `/` + v + ".json") {
				fmt.Println(v)
			}
		}
	}
	if f.Runlistjava {
		s := internal.Find(launcher.Minecraft + `/runtimes`)
		for _, v := range s {
			if internal.HasPrefixInSlice(v, internal.JavaRuntimeTypeNames) {
				fmt.Println(v)
			}
		}
	}
	if f.Download != "" {
		f.Outmsg = false
		f.D()
	}
	if f.Downloadfabric != "" {
		f.Outmsg = false
		f.Dfabric()
	}
	if f.Downloadquilt != "" {
		f.Outmsg = false
		f.Dquilt()
	}
	if f.Downloadpaper != "" {
		f.Outmsg = false
		f.Dpaper()
	}
	if f.Downloadjava != "" {
		f.Outmsg = false
		f.Djava()
	}
	if list {
		f.Listname()
	}
	if f.APIAddress != "" {
		f.Authlib()
	} else {
		f.APIAddress = "https://sessionserver.mojang.com"
	}
	if remove {
		f.Remove(ms)
		return
	}
	if ms {
		f.MsLogin()
	} else {
		if f.Email != "" {
			f.Aonline()
		} else {
			f.UUID = aflag.UUIDgen(f.Name)
			f.AccessToken = f.UUID
		}
	}
	if f.Runflag != "" {
		s := strings.Split(f.Runflag, " ")
		f.Flag = s
	}
	f.Gameinfo.RAM = f.RAM
	if f.Run != "" {
		if f.Name == "" && f.Email == "" {
			fmt.Println(lang.Lang("nousername"))
		} else {
			f.Arun()
		}
	}
	if f.Runpaper != "" {
		f.Arunpaper()
	}
	if update {
		check()
	}
}

type up struct {
	Version string `json:"version"`
	Msg     string `json:"msg"`
}

func check() {
	version, err := checkByDNS()
	if err != nil {
		log.Println(err)
		return
	}
	b, err := base64.StdEncoding.DecodeString(version)
	if err != nil {
		log.Println(err)
		return
	}
	u := up{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(lang.Lang("checkupdateerr"))
		fmt.Println(err)
		return
	}
	s, err := semver.NewVersion(u.Version)
	if err != nil {
		fmt.Println(lang.Lang("checkupdateerr"))
		fmt.Println(err)
		return
	}
	if s.GreaterThan(semver.MustParse(launcher.Launcherversion)) {
		fmt.Println(lang.Lang("checkupdate"), u.Version)
		fmt.Println(lang.Lang("nowversion"), launcher.Launcherversion)
		fmt.Println(lang.Lang("updateinfo"))
		fmt.Println(u.Msg)
	}
}

var Errtxt = errors.New("LookupTXT err")

func checkByDNS() (string, error) {
	l, err := net.LookupTXT("gml.xmdhs.com")
	if err != nil {
		return "", fmt.Errorf("checkByDNS: %w", err)
	}
	if len(l) != 1 {
		return "", fmt.Errorf("checkByDNS: %w", Errtxt)
	}
	return l[0], nil
}

func version() {
	fmt.Println("gomclauncher-" + internal.Launcherversion + "-" + internal.GitHash)
	fmt.Println("Build date: " + buildDate)
	fmt.Println("Build on: " + buildOn)
	fmt.Println("Repository: https://github.com/xmdhs/gomclauncher")
}
