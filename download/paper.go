package download

import (
	"encoding/json"
	"os"

	"github.com/xmdhs/gomclauncher/launcher"
	acopy "github.com/xmdhs/gomclauncher/thirdparty/copy"
)

func Newpaperjson(ver, verpaper, apath string) error {
	name := ver + `-paper-` + verpaper
	acopy.Copy(apath+"/versions/"+ver+"/"+ver+".json", apath+"/servers/"+name+"/"+name+".json")
	launcherjson := launcher.LauncherjsonX115{}
	b, err := os.ReadFile(apath + "/servers/" + name + "/" + name + ".json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &launcherjson)
	if err != nil {
		return err
	}
	launcherjson.ID = name
	marshal, err := json.Marshal(launcherjson)
	if err != nil {
		return err
	}
	err = os.WriteFile(apath+"/servers/"+name+"/"+name+".json", marshal, 0777)
	if err != nil {
		return err
	}
	return nil
}
