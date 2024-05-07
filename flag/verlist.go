package flag

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Arunlist() []string {
	if f.Verlistfabric {
		l, _ := f.Arunfabriclist()
		return l
	} else if f.Verlistquilt {
		l, _ := f.Arunquiltlist()
		return l
	}

	l, err := download.Getversionlist(context.Background(), f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	m := make(map[string]struct{})
	for _, v := range l.Versions {
		m[v.Type] = struct{}{}
	}
	var ok bool
	for k := range m {
		if f.Verlist == k {
			ok = true
		}
	}
	if ok {
		var versions []string
		for _, v := range l.Versions {
			if v.Type == f.Verlist {
				versions = append(versions, v.ID)
			}
		}

		return versions
	} else {
		fmt.Println(lang.Lang("runlist"))
		for k := range m {
			fmt.Println(k)
		}
	}

	return nil
}

func (f *Flag) Arunfabriclist() ([]string, error) {
	_, err := semver.NewVersion(f.Verlist)
	if err != nil {
		return nil, err
	}

	if f.Verlist != "" {
		l, err := download.Getfabricversionlist(context.Background(), f.Verlist, f.Atype, func(s string) { fmt.Println(s) })
		errr(err)
		var versions []string
		for _, v := range l.Versions {
			versions = append(versions, v.Loader.Version)
		}

		return versions, nil
	}

	return nil, nil
}

func (f *Flag) Arunquiltlist() ([]string, error) {
	_, err := semver.NewVersion(f.Verlist)
	if err != nil {
		return nil, err
	}

	if f.Verlist != "" {
		l, err := download.Getquiltversionlist(context.Background(), f.Verlist, f.Atype, func(s string) { fmt.Println(s) })
		errr(err)
		var versions []string
		for _, v := range l.Versions {
			versions = append(versions, v.Loader.Version)
		}

		return versions, nil
	}

	return nil, nil
}
