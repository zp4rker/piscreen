package screens

import (
	"image"
	"piscreen/util"
	"piscreen/vars"
)

type Home struct{}

func (s Home) Id() string {
	return "home"
}

func (s Home) Render() image.Image {
	context := util.BaseScreen(true)
	return context.Image()
}

func (s Home) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}

	switch key {
	case "KEY3":
		util.OpenMainMenu()
	case "KEY1":
		vars.Display.ToggleSleep()
	}
}
