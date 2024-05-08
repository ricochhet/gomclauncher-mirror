package flag

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Dfabric() {
	cxt := context.TODO()
	var err error
	var ver string
	if f.Run != "" {
		_, err = os.Stat(launcher.Minecraft + "/versions/" + f.Run + "/" + f.Run + ".json")
		ver = f.Run
	} else {
		_, err = os.Stat(launcher.Minecraft + "/versions/" + f.Download + "/" + f.Download + ".json")
		ver = f.Download
	}
	if err != nil {
		panic(err)
	}
	l, err := download.Getfabricversionlist(cxt, ver, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = l.Downfabricjson(cxt, f.Downloadfabric, launcher.Minecraft, func(s string) { fmt.Println(s) })
	if !(f.Run != "" && err != nil && errors.Is(err, download.NoSuch)) {
		errr(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(lang.Lang("finish"))
}
