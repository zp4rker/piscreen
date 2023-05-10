package standby

import (
	"piscreen/util"
	"piscreen/vars"
	"time"
)

func Observe() {
	go func() {
		for vars.Running {
			if !vars.Display.IsAsleep() && time.Now().Sub(vars.LastActive).Seconds() >= 15 {
				util.Debug("standby delay reached\n")
				vars.Display.ToggleSleep()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
