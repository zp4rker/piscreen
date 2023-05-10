package main

import (
	"flag"
	"image/color"
	"piscreen/keys"
	"piscreen/screens"
	"piscreen/spi"
	"piscreen/vars"
)

func main() {
	flag.IntVar(&vars.ListenDelay, "listen-delay", 170, "delay between button press listen")

	vars.Display = spi.NewST7789()
	defer vars.Display.Close()

	vars.Running = true
	keys.StartKeyListeners()

	vars.Display.Clear(color.RGBA{})

	for vars.Running {
		if vars.Display.Asleep {
			continue
		}

		vars.Display.ShowImage(screens.CurrentScreen.Render())
	}
}
