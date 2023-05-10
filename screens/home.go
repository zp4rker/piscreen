package screens

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"golang.org/x/exp/slices"
	"image"
	"piscreen/util"
	"piscreen/vars"
	"strings"
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

	lines := 0

	if vm, err := mem.VirtualMemory(); err == nil {
		memString := fmt.Sprintf("Memory: %.0f%%", vm.UsedPercent)
		context.DrawStringAnchored(memString, 5, float64(5+lines*17), 0, 1.25)
		lines++
	}

	if du, err := disk.Usage("/"); err == nil {
		diskString := fmt.Sprintf("Storage: %.2f%%", du.UsedPercent)
		context.DrawStringAnchored(diskString, 5, float64(5+lines*17), 0, 1.25)
		lines++
	}

	if ifs, err := net.Interfaces(); err == nil {
		for _, iface := range ifs {
			if slices.Contains(iface.Flags, "up") {
				for _, a := range iface.Addrs {
					if strings.Contains(a.Addr, ".") && strings.Contains(a.Addr, "/") {
						s := fmt.Sprintf("IP(%v): %v", iface.Name, a.Addr[:len(a.Addr)-3])
						context.DrawStringAnchored(s, 5, float64(5+lines*17), 0, 1.25)
						lines++
					}
				}
			}
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
