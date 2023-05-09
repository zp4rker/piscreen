package main

import (
	"image/color"
	"piscreen/keys"
	"piscreen/screens"
	"piscreen/spi"
	"piscreen/vars"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	vars.Running = true
	keys.StartKeyListeners()

	disp.Clear(color.RGBA{})

	for vars.Running {
		disp.ShowImage(screens.CurrentScreen.Render())
	}
}
