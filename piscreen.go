package main

import (
	"bytes"
	"github.com/fogleman/gg"
	"golang.org/x/image/bmp"
	"image/color"
	"piscreen/spi/impl"
)

func main() {
	println(byte(impl.RGBATo565(color.RGBA{R: 0xFF, A: 0xFF})))

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

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, context.Image()); err != nil {
		panic(err)
	}

	println(buf.Bytes()[0])
}
