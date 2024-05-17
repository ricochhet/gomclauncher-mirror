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

type Flag struct {
	launcher.Gameinfo
	Atype          string
	Downint        int
	Password       string
	Email          string
	Download       string
	Downloadserver bool
	Downloadfabric string
	Downloadquilt  string
	Downloadpaper  string
	Verlist        string
	Verlistfabric  string
	Verlistquilt   string
	Verlistpaper   string
	Run            string
	Runpaper       string
	Runlist        bool
	Runlistpaper   bool
	Runram         string
	Runflag        string
	Proxy          string
	Independent    bool
	Outmsg         bool
	Gmlconfig      Gmlconfig
	Saveconfig     func(gmlconfig Gmlconfig)
}

func NewFlag() *Flag {
	Gmlconfig := make(Gmlconfig)
	return &Flag{
		Gmlconfig:  Gmlconfig,
		Saveconfig: saveconfig,
	}
}

func (f *Flag) D() {
	cxt := context.TODO()
	l, err := download.Getversionlist(cxt, f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	err = l.Downjson(cxt, f.Download, launcher.Minecraft, func(s string) { fmt.Println(s) })
	if !(f.Run != "" && err != nil && errors.Is(err, download.ErrNoSuch)) {
		errr(err)
	}
	var b []byte
	if f.Run != "" {
		_, err := semver.NewVersion(f.Run)
		if err != nil {
			panic(err)
		}
		b, err = os.ReadFile(launcher.Minecraft + "/versions/" + f.Run + "/" + f.Run + ".json")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := semver.NewVersion(f.Download)
		if err != nil {
			panic(err)
		}
		b, err = os.ReadFile(launcher.Minecraft + "/versions/" + f.Download + "/" + f.Download + ".json")
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	dl, err := download.Newlibraries(cxt, b, f.Atype, func(s string) { fmt.Println(s) }, launcher.Minecraft)
	errr(err)
	if f.Outmsg {
		fmt.Println(lang.Lang("verifygamejar"))
	} else {
		fmt.Println(lang.Lang("downloadgamejar"))
	}
	err = dl.Downjar(f.Download)
	errr(err)
	err = dl.Downmappings(f.Download)
	errr(err)
	if f.Downloadserver {
		err = dl.Downserverjar(f.Download)
		errr(err)
		err = dl.Downservermappings(f.Download)
		errr(err)
	}
	fmt.Println(lang.Lang("finish"))
	if f.Outmsg {
		fmt.Println(lang.Lang("verifylibrarie"))
	} else {
		fmt.Println(lang.Lang("downloadlibrarie"))
	}
	f.dd(dl.Downlibrarie)
	fmt.Println(lang.Lang("finish"))
	if f.Outmsg {
		fmt.Println(lang.Lang("verifyassets"))
	} else {
		fmt.Println(lang.Lang("downloadassets"))
	}
	f.dd(dl.Downassets)
	fmt.Println(lang.Lang("finish"))
	if f.Outmsg {
		fmt.Println(lang.Lang("verifynatives"))
	} else {
		fmt.Println(lang.Lang("downloadnatives"))
	}
	err = dl.Unzip(f.Downint)
	if err != nil {
		fmt.Println(lang.Lang("downloadfail"))
		panic(err)
	}
	fmt.Println(lang.Lang("finish"))
}

func (f *Flag) dd(down func(i int, c chan int) error) {
	ch := make(chan int, 5)
	e := make(chan error)
	var err error
	go func() {
		err = down(f.Downint, ch)
		if err != nil {
			e <- err
		}
	}()
b:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				break b
			}
			if !f.Outmsg {
				fmt.Println(i)
			}
		case err := <-e:
			errr(err)
		}
	}
}

func errr(err error) {
	if err != nil {
		switch {
		case errors.Is(err, download.ErrNoSuch):
			fmt.Println(lang.Lang("download.NoSuch"))
		case errors.Is(err, download.ErrFileDownLoadFail):
			fmt.Println(lang.Lang("download.FileDownLoadFail"))
		default:
			panic(err)
		}
		os.Exit(0)
	}
}
