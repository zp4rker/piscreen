package standby

import (
	"piscreen/vars"
	"time"
)

var LastActive = time.Now()

func Observe() {
	go func() {
		for vars.Running {
			if !vars.Display.Asleep && time.Now().Sub(LastActive).Seconds() >= 15 {
				vars.Display.ToggleSleep()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
