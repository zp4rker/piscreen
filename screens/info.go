package screens

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"golang.org/x/exp/slices"
	"image"
	"piscreen/screens/menus"
	"piscreen/util"
	"strings"
)

type Info struct{}

func (s Info) Id() string {
	return "info"
}

func (s Info) Render() image.Image {
	context := BaseScreen(true)

	lines := 0

	if inf, err := host.Info(); err == nil {
		context.DrawStringAnchored(inf.Hostname, 5, float64(5+lines*17), 0, 1.25)
		lines++
	}

	if ifs, err := net.Interfaces(); err == nil {
		for _, iface := range ifs {
			if iface.Name == "lo" {
				// Skip lo
				continue
			}
			if slices.Contains(iface.Flags, "up") {
				for _, a := range iface.Addrs {
					if strings.Contains(a.Addr, ".") && strings.Contains(a.Addr, "/") {
						s := fmt.Sprintf("%v: %v", iface.Name, a.Addr[:len(a.Addr)-3])
						context.DrawStringAnchored(s, 5, float64(5+lines*17), 0, 1.25)
						lines++
					}
				}
			}
		}
		lines++
	}

	if cpup, err := cpu.Percent(0, false); err == nil {
		s := fmt.Sprintf("CPU: %.0f%%", cpup[0])
		context.DrawStringAnchored(s, 5, float64(5+lines*17), 0, 1.25)
		lines++
	}

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

	return context.Image()
}

func (s Info) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}

	switch key {
	case "KEY3":
		ChangeScreen(menus.Main{})
	}
}
