package main

import (
	"bytes"
	"github.com/fogleman/gg"
	"golang.org/x/image/bmp"
	"piscreen/spi/impl"
	"time"
)

func main() {
	disp := impl.NewST7789()
	defer disp.Close()

	context := gg.NewContext(240, 240)
	context.SetRGB(255, 0, 0)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, context.Image()); err != nil {
		panic(err)
	}

	println("Anything happen?")
	time.Sleep(5 * time.Second)
	println("How about now?")
}
