package screens

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"image"
	"piscreen/util"
	"piscreen/vars"
	"time"
)

type Home struct{}

func (s Home) Render() image.Image {
	context := Background()

	if err := context.LoadFontFace("JetBrainsMono.ttf", 17); err != nil {
		fmt.Printf("Failed to load font!")
	}

	// Footer
	context.SetRGB(util.GGColor(0xFF, 0xFF, 0xFF))
	context.DrawLine(0, 210, 240, 210)
	context.Stroke()
	now := time.Now()
	timeFmt := "Mon 2 Jan 2006 15:04"
	if now.Unix()%2 == 0 {
		timeFmt = "Mon 2 Jan 2006 15 04"
	}
	context.DrawStringAnchored(now.Format(timeFmt), 120, 215, 0.5, 1.25)

	vm, _ := mem.VirtualMemory()
	context.DrawString(fmt.Sprintf("Memory: %v%", vm.UsedPercent), 10, 10)

	return context.Image()
}

func (s Home) Handle(key string) {
	switch key {
	case "KEY_PRESS":
		vars.Display.ToggleSleep()
	}
}
