package main

import (
	"flag"
	"fmt"
	"image/color"
	"piscreen/keys"
	"piscreen/screens"
	"piscreen/spi"
	"piscreen/standby"
	"piscreen/util"
	"piscreen/vars"
)

func main() {
	flag.IntVar(&vars.ListenDelay, "listen-delay", 300, "delay between button press listen")
	flag.BoolVar(&vars.Debug, "debug", false, "whether to print debug messages")
	flag.Parse()
	util.Debug(fmt.Sprintf("listen-delay = %vms\n\n", vars.ListenDelay))

	vars.Display = spi.NewST7789()
	defer vars.Display.Close()
	util.Debug("display initialised\n")

	vars.Running = true
	keys.Listen()
	standby.Observe()
	util.Debug("go routines initialised\n\n")

	vars.Display.Clear(color.RGBA{})

	prevImage := screens.CurrentScreen.Render()

	for vars.Running {
		if vars.Display.Asleep {
			continue
		}

		newImage := screens.CurrentScreen.Render()
		if util.ImageCmp(prevImage, newImage) {
			continue
		}

		prevImage = newImage
		vars.Display.ShowImage(newImage)
		util.Debug("\nnew image rendered\n")
	}
}
