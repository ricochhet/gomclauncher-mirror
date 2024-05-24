package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

func Getarch() string {
	Arch := runtime.GOARCH
	switch Arch {
	case "amd64":
		return "64"
	case "386":
		return "32"
	default:
		panic("???")
	}
}

var ua string

func init() {
	b, _ := debug.ReadBuildInfo()
	var hash string
	for _, v := range b.Settings {
		if v.Key == "vcs.revision" {
			hash = v.Value
		}
	}
	ua = fmt.Sprintf("gomclauncher/%s (%v)", Launcherversion, hash)
	GitHash = hash
}

func HTTPGet(cxt context.Context, aurl string, t *http.Transport, header http.Header) (*http.Response, *time.Timer, error) {
	ctx, cancel := context.WithCancel(cxt)
	rep, err := http.NewRequestWithContext(ctx, http.MethodGet, aurl, nil)
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	if err != nil {
		return nil, nil, fmt.Errorf("HTTPGet: %w", err)
	}
	if header != nil {
		rep.Header = header
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", ua)
	c := http.Client{
		Transport: t,
	}
	reps, err := c.Do(rep)
	if err != nil {
		return reps, nil, fmt.Errorf("HTTPGet: %w", err)
	}
	return reps, timer, nil
}

var ErrPathInvalid = errors.New("path invalid")

func SafePathJoin(base string, path ...string) (string, error) {
	p := filepath.Join(append([]string{base}, path...)...)
	a, err := filepath.Rel(base, p)
	if err != nil {
		return "", fmt.Errorf("SafePathJoin: %w", err)
	}
	if strings.HasPrefix(a, ".") {
		return "", ErrPathInvalid
	}
	return p, nil
}

func Find(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	s := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			s = append(s, f.Name())
		}
	}
	return s
}

func SafeFind(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	s := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			s = append(s, f.Name())
		}
	}
	return s, nil
}

var JavaRuntimeTypeNames = []string{"alpha", "beta", "delta", "gamma", "gamma-snapshot", "jre-legacy", "minecraft-java-exe"}

func HasPrefixInSlice(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
