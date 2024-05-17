package flag

import (
	"fmt"

	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Remove(ms bool) {
	APIAddress := f.APIAddress
	if ms {
		APIAddress = "ms"
	}
	if f.Email == "" {
		fmt.Println(lang.Lang("emailnil"))
	}
	if _, ok := f.Gmlconfig[APIAddress][f.Email]; !ok {
		fmt.Println(APIAddress, f.Email, lang.Lang("nofind"))
	} else {
		delete(f.Gmlconfig[APIAddress], f.Email)
		fmt.Println(lang.Lang("removeok"), APIAddress, f.Email)
		if len(f.Gmlconfig[APIAddress]) == 0 {
			delete(f.Gmlconfig, APIAddress)
		}
		saveconfig(f.Gmlconfig)
	}
}
