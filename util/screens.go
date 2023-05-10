package util

import (
	"piscreen/core"
	"piscreen/vars"
)

func ChangeScreen(s core.Screen) {
	vars.PrevScreen = vars.CurrentScreen
	vars.CurrentScreen = s
}

func GoBackScreen() {
	s := vars.CurrentScreen
	vars.CurrentScreen = vars.PrevScreen
	vars.PrevScreen = s
}
