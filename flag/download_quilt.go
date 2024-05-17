package flag

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Dquilt() {
	cxt := context.TODO()
	var err error
	var ver string
	if f.Run != "" {
		_, err := semver.NewVersion(f.Run)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(launcher.Minecraft + "/versions/" + f.Run + "/" + f.Run + ".json")
		if err != nil {
			panic(err)
		}
		ver = f.Run
	} else {
		_, err := semver.NewVersion(f.Download)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(launcher.Minecraft + "/versions/" + f.Download + "/" + f.Download + ".json")
		if err != nil {
			panic(err)
		}
		ver = f.Download
	}
	l, err := download.Getquiltversionlist(cxt, ver, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = l.Downquiltjson(cxt, f.Downloadquilt, launcher.Minecraft, func(s string) { fmt.Println(s) })
	if !(f.Run != "" && err != nil && errors.Is(err, download.ErrNoSuch)) {
		errr(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(lang.Lang("finish"))
}
