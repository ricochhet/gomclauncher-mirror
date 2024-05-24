package download

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"runtime"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Javaruntimesjson launcher.Javaruntimesjson

type Javaruntimemanifest struct {
	t       string
	version string
	JSON    launcher.Javafilesjson
}

func Getjavaversionlist(cxt context.Context, atype string, print func(string)) (*Javaruntimesjson, error) {
	var b []byte
	r := newrandurls(atype)
	_, f := r.auto()

	err := retry.Do(func() error {
		url := `https://launchermeta.mojang.com/v1/products/java-runtime/2ec0cc96c44e5a76b9c8b7c39df7210883d12871/all.json`
		rep, _, err := Aget(cxt, url)
		if rep != nil {
			defer func() {
				if err := rep.Body.Close(); err != nil {
					panic(err)
				}
			}()
		}
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getjavaversionlistfail"), err, url)
		}
		b, err = io.ReadAll(rep.Body)
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getjavaversionlistfail"), err, url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return nil, fmt.Errorf("Getjavaversionlist: %w %w", err, ErrFileDownLoadFail)
	}

	v := Javaruntimesjson{}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, fmt.Errorf("Getjavaversionlist: %w", err)
	}
	return &v, nil
}

func (v Javaruntimesjson) Getjavaruntimemanifest(cxt context.Context, version string, atype string, print func(string)) (*Javaruntimemanifest, error) {
	var b []byte
	r := newrandurls(atype)
	_, f := r.auto()

	vv, err := v.os()
	if err != nil {
		return &Javaruntimemanifest{}, err
	}

	jrt, err := javaruntime(version, vv)
	if err != nil {
		return &Javaruntimemanifest{}, err
	}

	if len(jrt) == 0 {
		return &Javaruntimemanifest{}, fmt.Errorf("invalid java runtime version")
	}

	err = retry.Do(func() error {
		url := jrt[0].Manifest.URL
		rep, _, err := Aget(cxt, url)
		if rep != nil {
			defer func() {
				if err := rep.Body.Close(); err != nil {
					panic(err)
				}
			}()
		}
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getjavaruntimemanifestfail"), err, url)
		}
		b, err = io.ReadAll(rep.Body)
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getjavaruntimemanifestfail"), err, url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return nil, fmt.Errorf("Getjavaruntimemanifest: %w %w", err, ErrFileDownLoadFail)
	}

	ff := Javaruntimemanifest{}
	err = json.Unmarshal(b, &ff.JSON)
	ff.t = version
	ff.version = jrt[0].Version.Name
	if err != nil {
		return nil, fmt.Errorf("Getjavaruntimemanifest: %w", err)
	}
	return &ff, nil
}

func (v Javaruntimemanifest) Downjava(cxt context.Context, atype string, apath string, print func(string)) error {
	r := newrandurls(atype)
	_, f := r.auto()

	if len(v.JSON.Files) == 0 {
		return ErrNoSuch
	}

	for k, vv := range v.JSON.Files {
		if vv.Type == "file" {
			id := v.t + `-runtime-` + v.version
			path, err := internal.SafePathJoin(apath, `runtimes`, id, k)
			if ver(path, vv.Downloads.Raw.Sha1) {
				continue
			}
			if err != nil {
				return fmt.Errorf("downjava: %w", err)
			}

			err = retry.Do(func() error {
				url := vv.Downloads.Raw.URL
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
				return fmt.Errorf("downjava: %w %w", err, ErrFileDownLoadFail)
			}
		}
	}
	return nil
}

func (v Javaruntimesjson) os() (launcher.JavaRuntimeTypes, error) {
	os := runtime.GOOS
	arch := runtime.GOARCH

	switch {
	case os == "linux" && arch == "amd64":
		return v.Linux, nil
	case os == "linux" && arch == "386":
		return v.LinuxI386, nil
	case os == "darwin" && arch == "amd64":
		return v.MacOs, nil
	case os == "darwin" && arch == "arm64":
		return v.MacOsArm64, nil
	case os == "windows" && arch == "arm64":
		return v.WindowsArm64, nil
	case os == "windows" && arch == "amd64":
		return v.WindowsX64, nil
	case os == "windows" && arch == "386":
		return v.WindowsX86, nil
	default:
		return launcher.JavaRuntimeTypes{}, fmt.Errorf("unknown operating system or architecture")
	}
}

func javaruntime(version string, v launcher.JavaRuntimeTypes) ([]launcher.JavaRuntime, error) {
	switch {
	case version == "alpha":
		return v.JavaRuntimeAlpha, nil
	case version == "beta":
		return v.JavaRuntimeBeta, nil
	case version == "delta":
		return v.JavaRuntimeDelta, nil
	case version == "gamma":
		return v.JavaRuntimeGamma, nil
	case version == "gamma-snapshot":
		return v.JavaRuntimeGammaSnapshot, nil
	case version == "jre-legacy":
		return v.JreLegacy, nil
	case version == "minecraft-java-exe":
		return v.MinecraftJavaExe, nil
	default:
		return nil, fmt.Errorf("unknown java runtime version")
	}
}
