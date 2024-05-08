package flag

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Dpaper() {
	cxt := context.TODO()
	var err error
	var ver string
	var verpaper string
	if f.Runpaper != "" {
		ver = strings.Split(f.Runpaper, "-")[0]
		verpaper = strings.Split(f.Runpaper, "-")[1]
	} else {
		ver = strings.Split(f.Downloadpaper, "-")[0]
		verpaper = strings.Split(f.Downloadpaper, "-")[1]
	}
	dlver, err := download.Getversionlist(cxt, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = dlver.Downjson(cxt, ver, launcher.Minecraft, func(s string) { fmt.Println(s) })
	if !(f.Runpaper != "" && err != nil && errors.Is(err, download.NoSuch)) {
		errr(err)
	}
	_, err = os.Stat(launcher.Minecraft + "/versions/" + ver + "/" + ver + ".json")
	if err != nil || errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
	dlpaper, err := download.Getpaperversionlist(cxt, ver, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = dlpaper.Downpaperjar(cxt, verpaper, launcher.Minecraft, func(s string) { fmt.Println(s) })
	if !(f.Run != "" && err != nil && errors.Is(err, download.NoSuch)) {
		errr(err)
	}
	if err != nil {
		panic(err)
	}
	err = download.Newpaperjson(ver, verpaper, launcher.Minecraft)
	if err != nil {
		panic(err)
	}
	fmt.Println(lang.Lang("finish"))
}
