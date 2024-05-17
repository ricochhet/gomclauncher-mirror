package auth

import (
	"errors"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	a, err := Authenticate("https://sessionserver.mojang.com", "", "xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "")
	if !errors.Is(err, ErrNotOk) {
		t.Fatal(err)
	}
	t.Log(a)
}
