package screens

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
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
	memString := fmt.Sprintf("Memory: %d/%d MB", vm.Used/1024/1024, vm.Total/1024/1024)
	context.DrawStringAnchored(memString, 10, 10, 0, 1.25)

	du, _ := disk.Usage("/")
	diskString := fmt.Sprintf("Storage: %.2f/%.2f GB", float64(du.Used)/1024/1024/1024, float64(du.Total)/1024/1024/1024)
	context.DrawStringAnchored(diskString, 10, 27, 0, 1.25)

	return context.Image()
}

func (s Home) Handle(key string) {
	switch key {
	case "KEY_PRESS":
		vars.Display.ToggleSleep()
	}
}
