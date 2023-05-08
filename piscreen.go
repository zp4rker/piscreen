package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"piscreen/spi/impl"
)

func main() {
	println(byte(impl.RGBATo565(color.RGBA{R: 0xFF, A: 0xFF})))
	fmt.Printf("%+v\n", color.RGBA{R: 0x01, A: 0xFF})
	println(byte(impl.RGBATo565(color.RGBA{R: 0x01, A: 0xFF})))

	disp := impl.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear()
	println("done")

	context := gg.NewContext(240, 240)
	context.SetRGB(255, 0, 0)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	println("printing image...")
	disp.ShowImage(context.Image())
	println("done")

	fmt.Printf("%v\n", context.Image().At(0, 0))
	println(impl.RGBATo565(context.Image().At(0, 0)))
}
