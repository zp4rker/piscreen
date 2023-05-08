package main

import (
	"piscreen/screens"
	"piscreen/spi"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear()
	println("done")

	println("printing image...")
	disp.ShowImage(screens.CurrentScreen.Render())
	println("done")
}
