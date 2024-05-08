package launcher

import (
	"errors"
	"fmt"
	"os"
)

func (g *Gameinfo) Runpaper115() (err error) {
	l, _, err := g.GenPaperCmdArgs()
	if err != nil {
		return fmt.Errorf("Runpaper115: %w", err)
	}
	err = l.Launcher115()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Runpaper115: %w", ErrJavaPath)
		}
		return fmt.Errorf("Runpaper115: %w", err)
	}
	return nil
}

func (g *Gameinfo) GenPaperCmdArgs() (l *launcher1155, args []string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()
	err = creatlauncherprofiles(g)
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenPaperCmdArgs: %w", err)
	}
	l, err = g.modjson()
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenPaperCmdArgs: %w", err)
	}
	l.Gamedir = g.Minecraftpath + `/servers/` + l.json.ID + `/`
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-Xms`+g.RAM+`m`)
	if g.Flag != nil {
		l.flag = append(l.flag, g.Flag...)
	}
	l.flag = append(l.flag, `-cp`)
	l.flag = append(l.flag, l.json.ID+`.jar`)
	l.flag = append(l.flag, `io.papermc.paperclip.Main`)
	l.flag = append(l.flag, `--nogui`)
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenPaperCmdArgs: %w", err)
	}
	return l, l.flag, nil
}
