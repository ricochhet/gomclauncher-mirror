package download

import (
	"encoding/json"
	"os"
)

type Paperid struct {
	ID string `json:"id"`
}

func Newpaperjson(ver, verpaper, apath string) error {
	name := ver + `-paper-` + verpaper
	b, err := json.Marshal(Paperid{ID: name})
	if err != nil {
		return err
	}
	err = os.WriteFile(apath+"/servers/"+name+"/"+name+".json", b, 0777)
	if err != nil {
		return err
	}
	return nil
}
