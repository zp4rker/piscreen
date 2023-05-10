package screens

import (
	"fmt"
	"github.com/fogleman/gg"
	"piscreen/util"
	"time"
)

func BaseScreen(footer bool) *gg.Context {
	context := gg.NewContext(240, 240)

	context.SetRGB(util.GGColor(0x35, 0x37, 0x39))
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	if err := context.LoadFontFace("JetBrainsMono.ttf", 17); err != nil {
		fmt.Printf("Failed to load font!")
	}

	if footer {
		context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
		context.DrawLine(0, 210, 240, 210)
		context.Stroke()
		now := time.Now()
		timeFmt := "Mon 2 Jan 2006 15:04"
		if now.Unix()%2 == 0 {
			timeFmt = "Mon 2 Jan 2006 15 04"
		}
		context.DrawStringAnchored(now.Format(timeFmt), 120, 215, 0.5, 1.25)
	}

	return context
}
