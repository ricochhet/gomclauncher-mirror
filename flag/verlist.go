package flag

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Masterminds/semver/v3"
	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Arunlist() []string {
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
	_, err := semver.NewVersion(f.Verlistfabric)
	if err != nil {
		return nil, err
	}

	if f.Verlistfabric != "" {
		l, err := download.Getfabricversionlist(context.Background(), f.Verlistfabric, f.Atype, func(s string) { fmt.Println(s) })
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
	_, err := semver.NewVersion(f.Verlistquilt)
	if err != nil {
		return nil, err
	}

	if f.Verlistquilt != "" {
		l, err := download.Getquiltversionlist(context.Background(), f.Verlistquilt, f.Atype, func(s string) { fmt.Println(s) })
		errr(err)
		var versions []string
		for _, v := range l.Versions {
			versions = append(versions, v.Loader.Version)
		}

		return versions, nil
	}

	return nil, nil
}

func (f *Flag) Arunpaperlist() ([]string, error) {
	_, err := semver.NewVersion(f.Verlistpaper)
	if err != nil {
		return nil, err
	}

	if f.Verlistpaper != "" {
		l, err := download.Getpaperversionlist(context.Background(), f.Verlistpaper, f.Atype, func(s string) { fmt.Println(s) })
		errr(err)
		var versions []string
		for _, v := range l.Json.Builds {
			versions = append(versions, strconv.Itoa(v.Build))
		}

		return versions, nil
	}

	return nil, nil
}
