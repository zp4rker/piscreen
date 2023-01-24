package main

import (
	"github.com/fogleman/gg"
	"github.com/stianeikeland/go-rpio/v4"
	"piscreen/spi/impl"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	display := impl.NewST7789()
	context := gg.NewContext(240, 240)

	context.SetRGB(255, 255, 255)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	display.ShowImage(context.Image())
}
