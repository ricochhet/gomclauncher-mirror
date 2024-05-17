package download

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/launcher"
)

func TestNewlibrarie(t *testing.T) {
	b, err := os.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "", func(s string) { fmt.Println(s) }, launcher.Minecraft)
	if err != nil {
		t.Fatal(err)
	}
	if l.librarie.ID != "1.15.2" {
		t.Fatal(l.librarie.ID)
	}
}

func TestAssets(t *testing.T) {
	bb, err := os.ReadFile("1.16.json")
	if err != nil {
		t.Fatal(err)
	}
	a := assets{}
	err = json.Unmarshal(bb, &a)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	auth.Transport = http.DefaultTransport.(*http.Transport).Clone()
	code := m.Run()
	os.Exit(code)
}

func Test_get(t *testing.T) {
	err := get(context.Background(), "https://piston-meta.mojang.com/mc/game/version_manifest.json", "test/test/a.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = os.Stat("test/test/a.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	if err := os.RemoveAll("test"); err != nil {
		t.Fatal(err)
	}
}

func Test_source(t *testing.T) {
	type args struct {
		url   string
		types string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				url:   "https://piston-meta.mojang.com/mc/game/version_manifest.json",
				types: "bmclapi",
			},
			want: "https://bmclapi2.bangbang93.com/mc/game/version_manifest.json",
		},
		{
			name: "2",
			args: args{
				url:   "https://piston-meta.mojang.com/mc/game/version_manifest.json",
				types: "vvv",
			},
			want: "https://piston-meta.mojang.com/mc/game/version_manifest.json",
		},
		{
			name: "3",
			args: args{
				url:   "https://piston-meta.mojang.com/mc/game/version_manifest.json/piston-meta.mojang.com",
				types: "bmclapi",
			},
			want: "https://bmclapi2.bangbang93.com/mc/game/version_manifest.json/piston-meta.mojang.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := source(tt.args.url, tt.args.types); got != tt.want {
				t.Errorf("source() = %v, want %v", got, tt.want)
			}
		})
	}
}
