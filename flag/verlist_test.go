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
	f.Verlist = "1.20.4"
	versionsfabric, err := f.Arunfabriclist()
	if err != nil {
		t.Fatal(err)
	}

	if len(versionsfabric) == 0 {
		t.Fatal(fmt.Errorf("versionsfabric is null or empty"))
	}
}
