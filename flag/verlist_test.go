package flag

import (
	"fmt"
	"testing"
)

var f *Flag = NewFlag()

func TestVerlist(t *testing.T) {
	f.Verlist = "release"
	versions := f.Arunlist()
	if len(versions) == 0 {
		t.Fatal(fmt.Errorf("versions is null or empty"))
	}
}

func TestVerlistfabric(t *testing.T) {
	f.Verlistfabric = "1.20.4"
	versionsfabric, err := f.Arunfabriclist()
	if err != nil {
		t.Fatal(err)
	}

	if len(versionsfabric) == 0 {
		t.Fatal(fmt.Errorf("versionsfabric is null or empty"))
	}
}

func TestVerlistquilt(t *testing.T) {
	f.Verlistquilt = "1.20.4"
	versionsquilt, err := f.Arunquiltlist()
	if err != nil {
		t.Fatal(err)
	}

	if len(versionsquilt) == 0 {
		t.Fatal(fmt.Errorf("versionsquilt is null or empty"))
	}
}

func TestVerlistpaper(t *testing.T) {
	f.Verlistpaper = "1.20.4"
	versionspaper, err := f.Arunpaperlist()
	if err != nil {
		t.Fatal(err)
	}

	if len(versionspaper) == 0 {
		t.Fatal(fmt.Errorf("versionspaper is null or empty"))
	}
}

func TestVerlistjava(t *testing.T) {
	versionsjava, err := f.Arunjavalist()
	if err != nil {
		t.Fatal(err)
	}

	if len(versionsjava) == 0 {
		t.Fatal(fmt.Errorf("versionsjava is null or empty"))
	}
}
