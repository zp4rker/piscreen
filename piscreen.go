package main

import (
	"github.com/fogleman/gg"
	"piscreen/spi/impl"
)

func main() {
	disp := impl.NewST7789()
	defer disp.Close()

	disp.Clear()

	context := gg.NewContext(240, 240)
	context.SetRGB(255, 0, 0)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	disp.ShowImage(context.Image())
}
