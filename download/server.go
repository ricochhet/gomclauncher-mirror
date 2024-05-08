package download

import (
	"encoding/json"
	"os"
)

type Serverid struct {
	ID string `json:"id"`
}

func Newserverjson(ver, apath string) error {
	err := serverjson(ver, apath)
	return err
}

func Newpaperjson(ver, verpaper, apath string) error {
	name := ver + `-paper-` + verpaper
	err := serverjson(name, apath)
	return err
}

func serverjson(name, apath string) error {
	b, err := json.Marshal(Serverid{ID: name})
	if err != nil {
		return err
	}
	err = os.WriteFile(apath+"/servers/"+name+"/"+name+".json", b, 0777)
	if err != nil {
		return err
	}
	return nil
}
