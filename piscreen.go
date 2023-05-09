package main

import (
	"image/color"
	"piscreen/keys"
	"piscreen/screens"
	"piscreen/spi"
	"piscreen/vars"
)

func main() {
	vars.Display = spi.NewST7789()
	defer vars.Display.Close()

	vars.Running = true
	keys.StartKeyListeners()

	vars.Display.Clear(color.RGBA{})

	for vars.Running {
		if vars.Asleep {
			continue
		}

		vars.Display.ShowImage(screens.CurrentScreen.Render())
	}
}
