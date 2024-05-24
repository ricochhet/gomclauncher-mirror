package flag

import (
	"encoding/json"
	"os"
)

func Test(path string) bool {
	b, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	t := t{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return false
	}
	if len(t.Libraries) == 0 {
		return false
	}
	if t.MainClass == "" {
		return false
	}
	return true
}

func Testservers(path string) bool {
	b, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	t := t2{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return false
	}
	if t.ID == "" {
		return false
	}
	return true
}

type t struct {
	Libraries []interface{} `json:"Libraries"`
	MainClass string        `json:"mainClass"`
}

type t2 struct {
	ID string `json:"id"`
}
