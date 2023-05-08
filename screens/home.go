package screens

import (
	"image"
	"piscreen/util"
)

type Home struct{}

func (s Home) Render() image.Image {
	context := Background()

	context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
	context.DrawLine(0, 210, 240, 210)
	context.Stroke()

	return context.Image()
}
