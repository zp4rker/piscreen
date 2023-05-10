package util

import "piscreen/vars"

func Debug(msg string) {
	if !vars.Debug {
		return
	}
	println(msg)
}
