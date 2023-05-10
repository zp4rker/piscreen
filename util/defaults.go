package util

import (
	"piscreen/standby"
	"piscreen/vars"
	"time"
)

func DefaultHandle(_ string) bool {
	standby.LastActive = time.Now()
	if vars.Display.Asleep {
		vars.Display.ToggleSleep()
		return true
	}
	return false
}
