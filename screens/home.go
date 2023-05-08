package screens

import (
	"fmt"
	"image"
	"piscreen/util"
	"time"
)

type Home struct{}

func (s Home) Render() image.Image {
	context := Background()

	context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
	context.DrawLine(0, 210, 240, 210)
	context.Stroke()

	if err := context.LoadFontFace("JetBrainsMono.ttf", 17); err != nil {
		fmt.Printf("Failed to load font!")
	}
	context.DrawStringAnchored(time.Now().Format("Mon 2 Jan 2006 15:04"), 120, 215, 0.5, 0)

	return context.Image()
}
