package screens

import (
	"github.com/shirou/gopsutil/v3/host"
	"image"
	"piscreen/util"
	"strconv"
)

type Home struct{}

func (s Home) Id() string {
	return "home"
}

func (s Home) Render() image.Image {
	context := util.BaseScreen(true)

	lines := 0
	if inf, err := host.Uptime(); err == nil {
		util.White(context)
		context.DrawStringAnchored("Uptime", 5, float64(5+lines*17), 0, 1.25)
		lines++
		util.Grey(context)
		context.DrawStringAnchored(util.TimeString(inf), 5, float64(5+lines*17), 0, 1.25)
		lines++
	}

	if inf, err := host.SensorsTemperatures(); err == nil {
		if lines > 0 {
			lines++
		}
		for _, temp := range inf {
			util.White(context)
			context.DrawStringAnchored("CPU Temp", 5, float64(5+lines*17), 0, 1.25)
			lines++
			util.Grey(context)
			context.DrawStringAnchored(strconv.FormatFloat(temp.Temperature, 'f', 1, 32), 5, float64(5+lines*17), 0, 1.25)
			lines++
		}
	}

	return context.Image()
}

func (s Home) Handle(key string) {
	if util.DefaultHandle(key) {
		return
	}
}
