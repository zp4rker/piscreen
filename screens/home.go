package screens

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"image"
	"piscreen/util"
	"piscreen/vars"
	"time"
)

type Home struct{}

func (s Home) Render() image.Image {
	context := Background()

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

	if vm, err := mem.VirtualMemory(); err == nil {
		memString := fmt.Sprintf("Memory: %.0f%%", vm.UsedPercent)
		context.DrawStringAnchored(memString, 5, 5, 0, 1.25)
	}

	if du, err := disk.Usage("/"); err == nil {
		diskString := fmt.Sprintf("Storage: %.2f%%", du.UsedPercent)
		context.DrawStringAnchored(diskString, 5, 26, 0, 1.25)
	}

	if ifs, err := net.Interfaces(); err == nil {
		for _, iface := range ifs {
			println(iface.Name)
			fmt.Printf("%v\n", iface.Flags)
			fmt.Printf("%v\n", iface.Addrs)
		}
	}

	return context.Image()
}

func (s Home) Handle(key string) {
	switch key {
	case "KEY_PRESS":
		vars.Display.ToggleSleep()
	}
}
