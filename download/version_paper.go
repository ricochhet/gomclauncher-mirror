package download

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Paperjsonv2 struct {
	atype   string
	version string
	Json    launcher.Paperjsonv2
}

func Getpaperversionlist(cxt context.Context, version, atype string, print func(string)) (*Paperjsonv2, error) {
	var b []byte
	r := newrandurls(atype)
	_, f := r.auto()

	err := retry.Do(func() error {
		url := source(`https://api.papermc.io/v2/projects/paper/versions/`+version+`/builds`, f)
		rep, _, err := Aget(cxt, url)
		if rep != nil {
			defer rep.Body.Close()
		}
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getpaperversionlistfail"), err, url)
		}
		b, err = io.ReadAll(rep.Body)
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getpaperversionlistfail"), err, url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return nil, fmt.Errorf("Getpaperversionlist: %w %w", err, FileDownLoadFail)
	}

	v := Paperjsonv2{}
	err = json.Unmarshal(b, &v.Json)
	v.atype = atype
	v.version = version
	if err != nil {
		return nil, fmt.Errorf("Getpaperversionlist: %w", err)
	}
	return &v, nil
}

func (v Paperjsonv2) Downpaperjar(cxt context.Context, version, apath string, print func(string)) error {
	r := newrandurls(v.atype)
	_, f := r.auto()
	for _, vv := range v.Json.Builds {
		if strconv.Itoa(vv.Build) == version {
			id := v.version + `-paper-` + version
			path, err := internal.SafePathJoin(apath, `/servers/`, id, id+".jar")
			if ver(path, vv.Downloads.Application.SHA256) {
				return nil
			}
			if err != nil {
				return fmt.Errorf("Downpaperjson: %w", err)
			}

			err = retry.Do(func() error {
				url := source(`https://api.papermc.io/v2/projects/paper/versions/`+v.version+`/builds/`+version+`/downloads/`+`paper-`+v.version+`-`+version+`.jar`, f)
				err := get(cxt, url, path)
				if err != nil {
					f = r.fail(f)
					return fmt.Errorf("%v %v %w", lang.Lang("weberr"), url, err)
				}
				return nil
			}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
				print(fmt.Sprintf("retry %d: %v", n, err))
			}))...)
			if err != nil {
				return fmt.Errorf("Downpaperjson: %w %w", err, FileDownLoadFail)
			}
			err = Newpaperjson(v.version, version, launcher.Minecraft)
			if err != nil {
				panic(err)
			}
			return nil
		}
	}
	return NoSuch
}
