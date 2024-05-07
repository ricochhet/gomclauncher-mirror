package download

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Fabricjsonv2 struct {
	atype    string
	version  string
	Versions []launcher.Fabricjsonv2
}

func Getfabricversionlist(cxt context.Context, version, atype string, print func(string)) (*Fabricjsonv2, error) {
	var b []byte
	r := newrandurls(atype)
	_, f := r.auto()

	err := retry.Do(func() error {
		url := source(`https://meta.fabricmc.net/v2/versions/loader/`+version, f)
		rep, _, err := Aget(cxt, url)
		if rep != nil {
			defer rep.Body.Close()
		}
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getfabricversionlistfail"), err, url)
		}
		b, err = io.ReadAll(rep.Body)
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getfabricversionlistfail"), err, url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return nil, fmt.Errorf("Getfabricversionlist: %w %w", err, FileDownLoadFail)
	}

	v := Fabricjsonv2{}
	err = json.Unmarshal(b, &v.Versions)
	v.atype = atype
	v.version = version
	if err != nil {
		return nil, fmt.Errorf("Getfabricversionlist: %w", err)
	}
	return &v, nil
}

func (v Fabricjsonv2) Downfabricjson(cxt context.Context, version, apath string, print func(string)) error {
	r := newrandurls(v.atype)
	_, f := r.auto()
	for _, vv := range v.Versions {
		if vv.Loader.Version == version {
			id := v.version + `-fabric-` + version
			path, err := internal.SafePathJoin(apath, `/versions/`, id, id+".json")
			if err != nil {
				return fmt.Errorf("Downfabricjson: %w", err)
			}

			err = retry.Do(func() error {
				url := source(`https://meta.fabricmc.net/v2/versions/loader/`+v.version+`/`+version+`/profile/json`, f)
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
				return fmt.Errorf("Downfabricjson: %w %w", err, FileDownLoadFail)
			}
			return nil
		}
	}
	return NoSuch
}
