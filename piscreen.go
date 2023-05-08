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

	img := context.Image()
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			fmt.Printf("%+v\n", img.At(x, y))
		}
	}
}
