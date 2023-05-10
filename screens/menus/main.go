package menus

import (
	"image"
	"piscreen/util"
)

type Main struct{}

var (
	buttons = []string{
		"Home",
		"Something",
		"Something else",
		"Info",
		"Exit",
	}

	focus = 0
)

func (s Main) Id() string {
	return "main_menu"
}

func (s Main) Render() image.Image {
	context := util.BaseScreen(false)

	for i := 0; i < len(buttons); i++ {
		context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
		context.DrawRectangle(10, float64(10+i*46), 220, 36)
		if focus == i {
			context.Fill()
			context.SetRGB(util.GGColor(0x00, 0x00, 0x00))
		} else {
			context.Stroke()
		}

		context.DrawStringAnchored(buttons[i], 120, float64(10+i*46+18), 0.5, 0.5)
	}

	return context.Image()
}

func (s Main) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}

	switch key {
	case "KEY1":
		util.GoBackScreen()
	case "KEY_UP":
		if focus > 0 {
			focus--
		} else {
			focus = len(buttons) - 1
		}
	case "KEY_DOWN":
		if focus < len(buttons)-1 {
			focus++
		} else {
			focus = 0
		}
	}
}
