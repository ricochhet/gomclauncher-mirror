package flag

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Aonline() {
	if f.Email == "" {
		fmt.Println(lang.Lang("emailnil"))
		os.Exit(0)
	}
	if f.Gmlconfig[f.APIAddress] == nil {
		f.Gmlconfig[f.APIAddress] = make(map[string]Config)
	}
	err := f.Gmlconfig[f.APIAddress][f.Email].setonline(&f.Gmlconfig, f)
	if err != nil {
		switch err { //nolint:errorlint // ...
		case ErrHaveProfiles:
			a := auth.Auth{
				AccessToken: f.Gmlconfig[f.APIAddress][f.Email].AccessToken,
				ClientToken: f.Gmlconfig[f.APIAddress][f.Email].ClientToken,
				APIAddress:  f.APIAddress,
				Username:    f.Gmlconfig[f.APIAddress][f.Email].Name,
				ID:          f.Gmlconfig[f.APIAddress][f.Email].UUID,
			}
			atime := time.Now().Unix()
			if atime-f.Gmlconfig[f.APIAddress][f.Email].Time > 120 {
				if err := auth.Validate(a); err != nil {
					err = auth.Refresh(&a)
					if err != nil {
						if errors.Is(err, auth.ErrNotOk) {
							fmt.Println(lang.Lang("auth.NotOk-refresh"))
							os.Exit(0)
						} else {
							fmt.Println(lang.Lang("Refresherr"))
							log.Println(err)
							os.Exit(0)
						}
					}
					aconfig := f.Gmlconfig[f.APIAddress][f.Email]
					aconfig.Name = a.Username
					aconfig.UUID = a.ID
					aconfig.AccessToken = a.AccessToken
					aconfig.Time = time.Now().Unix()
					aconfig.ClientToken = a.ClientToken
					f.Gmlconfig[f.APIAddress][f.Email] = aconfig
					saveconfig(f.Gmlconfig)
				}
			}
		case auth.ErrNotOk:
			fmt.Println(lang.Lang("auth.NotOk"))
			os.Exit(0)
		default:
			panic(err)
		}
	}
	if f.Gmlconfig[f.APIAddress][f.Email].Name == "" {
		fmt.Println(lang.Lang("namenil"))
		os.Exit(0)
	}
	f.AccessToken = f.Gmlconfig[f.APIAddress][f.Email].AccessToken
	f.Name = f.Gmlconfig[f.APIAddress][f.Email].Name
	f.UUID = f.Gmlconfig[f.APIAddress][f.Email].UUID
}

func (f *Flag) Listname() {
	fmt.Println("-----------------")
	for k, v := range f.Gmlconfig {
		switch k {
		case "https://sessionserver.mojang.com":
			fmt.Println(lang.Lang("minecraftlogin"))
		case "ms":
			fmt.Println(lang.Lang("mslogin"))
		default:
			fmt.Println(lang.Lang("authlib-injectorlogin"), k)
		}
		for k, v := range v {
			fmt.Println(lang.Lang("email"), k)
			fmt.Println(lang.Lang("name"), v.Name)
		}
		fmt.Println("-----------------")
	}
}
