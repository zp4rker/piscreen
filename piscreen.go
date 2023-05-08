package main

import (
	"image/color"
	"piscreen/screens"
	"piscreen/spi"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear(color.RGBA{})
	println("done")

	println("printing image...")
	disp.ShowImage(screens.CurrentScreen.Render())
	println("done")
}
