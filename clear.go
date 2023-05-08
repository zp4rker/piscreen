package main

import (
	"image/color"
	"piscreen/spi"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear(color.RGBA{})
	println("done")
}
