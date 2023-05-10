package util

import (
	"piscreen/vars"
	"time"
)

func DefaultHandle(_ string) bool {
	vars.LastActive = time.Now()
	if vars.Display.IsAsleep() {
		vars.Display.ToggleSleep()
		return true
	}
	return false
}
