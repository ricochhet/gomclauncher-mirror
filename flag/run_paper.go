package flag

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Arunpaper() {
	f.Version = f.Runpaper
	if f.Independent {
		f.Gamedir = f.Minecraftpath + "/servers/" + f.Version
	} else {
		f.Gamedir = f.Minecraftpath
	}
	b, err := os.ReadFile(f.Minecraftpath + "/servers/" + f.Version + "/" + f.Version + ".json")
	if err != nil {
		fmt.Println(lang.Lang("nofindthisversion"))
		panic(err)
	}
	if f.Outmsg {
		t := map[string]interface{}{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			panic(err)
		}
		id, _ := t["id"].(string)
		if id != f.Version {
			t["id"] = f.Version
			b, err = json.Marshal(t)
			if err != nil {
				panic(err)
			}
			err := os.WriteFile(f.Minecraftpath+"/servers/"+f.Version+"/"+f.Version+".json", b, 0o777)
			if err != nil {
				panic(err)
			}
		}
	}
	f.Jsonbyte = b
	err = f.Runpaper115()
	if err != nil {
		var e launcher.ErrLegacyNoExit
		if errors.As(err, &e) {
			fmt.Println(lang.Lang("legacynoexit"), err)
			os.Exit(0)
		}
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(lang.Lang("flag.os.ErrNotExist"))
			os.Exit(0)
		}
		if errors.Is(err, launcher.ErrJSON) {
			fmt.Printf(lang.Lang("launcher.JsonErr"), launcher.Minecraft)
			os.Exit(0)
		}
		if errors.Is(err, launcher.ErrJSONNorTrue) {
			fmt.Println(lang.Lang("launcher.JsonNorTrue"))
			os.Exit(0)
		}
		panic(err)
	}
}
