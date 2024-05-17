package download

import (
	"fmt"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (l Libraries) Downserverjar(version string) error {
	path, err := internal.SafePathJoin(l.path, `/servers/`, version, version+".jar")
	if err != nil {
		return fmt.Errorf("Downserverjar: %w %w", err, ErrFileDownLoadFail)
	}
	if ver(path, l.librarie.Downloads.Server.Sha1) {
		return nil
	}
	_, t := l.auto()

	err = retry.Do(func() error {
		u := source(l.librarie.Downloads.Server.URL, t)
		err := get(l.cxt, u, path)
		if err != nil {
			t = l.fail(t)
			return fmt.Errorf("%v %w %v", lang.Lang("weberr"), err, u)
		}
		if !ver(path, l.librarie.Downloads.Server.Sha1) {
			t = l.fail(t)
			return fmt.Errorf("%v %v", lang.Lang("filecheckerr"), u)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		l.print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return fmt.Errorf("Downserverjar: %w %w", err, ErrFileDownLoadFail)
	}
	err = Newserverjson(version, launcher.Minecraft)
	if err != nil {
		panic(err)
	}
	return nil
}

func (l Libraries) Downservermappings(version string) error {
	path, err := internal.SafePathJoin(l.path, `/servers/`, version, version+"-mappings.txt")
	if err != nil {
		return fmt.Errorf("Downservermappings: %w %w", err, ErrFileDownLoadFail)
	}
	if ver(path, l.librarie.Downloads.ServerMappings.Sha1) {
		return nil
	}
	_, t := l.auto()

	err = retry.Do(func() error {
		u := source(l.librarie.Downloads.ServerMappings.URL, t)
		err := get(l.cxt, u, path)
		if err != nil {
			t = l.fail(t)
			return fmt.Errorf("%v %w %v", lang.Lang("weberr"), err, u)
		}
		if !ver(path, l.librarie.Downloads.ServerMappings.Sha1) {
			t = l.fail(t)
			return fmt.Errorf("%v %v", lang.Lang("filecheckerr"), u)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		l.print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return fmt.Errorf("Downservermappings: %w %w", err, ErrFileDownLoadFail)
	}
	return nil
}
