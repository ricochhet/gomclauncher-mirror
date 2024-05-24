package flag

import (
	"context"
	"errors"
	"fmt"

	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Djava() {
	cxt := context.TODO()
	var err error
	dljava, err := download.Getjavaversionlist(cxt, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	manifest, err := dljava.Getjavaruntimemanifest(cxt, f.Downloadjava, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = manifest.Downjava(cxt, f.Atype, launcher.Minecraft, func(s string) { fmt.Println(s) })
	errr(err)
	if !(f.Run != "" && err != nil && errors.Is(err, download.ErrNoSuch)) {
		errr(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(lang.Lang("finish"))
}
