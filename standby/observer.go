package standby

import (
	"piscreen/util"
	"piscreen/vars"
	"time"
)

var LastActive = time.Now()

func Observe() {
	go func() {
		for vars.Running {
			if !vars.Display.Asleep && time.Now().Sub(LastActive).Seconds() >= 15 {
				util.Debug("standby delay reached\n")
				vars.Display.ToggleSleep()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
