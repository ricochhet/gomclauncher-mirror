package download

import (
	"context"
	"fmt"
	"testing"

	"github.com/xmdhs/gomclauncher/launcher"
)

func TestGetversionlist(t *testing.T) {
	cxt := context.TODO()
	l, err := Getversionlist(cxt, "vanilla", func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
	err = l.Downjson(cxt, "1.20.4", launcher.Minecraft, func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetfabricversionlist(t *testing.T) {
	cxt := context.TODO()
	l, err := Getfabricversionlist(cxt, "1.20.4", "vanilla", func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
	err = l.Downfabricjson(cxt, "0.15.11", launcher.Minecraft, func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetquiltversionlist(t *testing.T) {
	cxt := context.TODO()
	l, err := Getquiltversionlist(cxt, "1.20.4", "vanilla", func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
	err = l.Downquiltjson(cxt, "0.25.0", launcher.Minecraft, func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
}
