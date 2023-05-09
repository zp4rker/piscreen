package screens

import (
	"github.com/fogleman/gg"
	"image"
	"piscreen/util"
)

var CurrentScreen Screen = Home{}

type Screen interface {
	Render() image.Image
	Handle(key string)
}

func Background() *gg.Context {
	context := gg.NewContext(240, 240)

	context.SetRGB(util.GGColor(0x35, 0x37, 0x39))
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	return context
}
