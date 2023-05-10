package main

import (
	"flag"
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

	vars.Display = spi.NewST7789()
	defer vars.Display.Close()

	vars.Running = true
	keys.Listen()
	standby.Observe()

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
	}
}
