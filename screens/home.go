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
	now := time.Now()
	timeFmt := "Mon 2 Jan 2006 15:04"
	if now.Unix()%2 == 0 {
		timeFmt = "Mon 2 Jan 2006 15 04"
	}
	context.DrawStringAnchored(now.Format(timeFmt), 120, 215, 0.5, 1.25)

	return context.Image()
}
