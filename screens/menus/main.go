package menus

import (
	"fmt"
	"image"
	"piscreen/screens"
	"piscreen/util"
)

type Main struct{}

func (s Main) Id() string {
	return "main_menu"
}

func (s Main) Render() image.Image {
	context := screens.BaseScreen(false)

	context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
	for i := 0; i < 5; i++ {
		context.DrawRectangle(10, float64(10+i*46), 220, 36)
		context.Stroke()

		context.DrawStringAnchored(fmt.Sprintf("Button %v", i), 120, float64(10+i*46+18), 0.5, 0.5)
	}

	return context.Image()
}

func (s Main) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}

	switch key {
	case "KEY1":
		screens.GoBack()
	}
}
