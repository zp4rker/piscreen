package main

import (
	"github.com/fogleman/gg"
	"github.com/manx98/go-st7789"
	"github.com/stianeikeland/go-rpio/v4"
	"piscreen/impl"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	device := ST7789.NewST7789(
		&impl.SPI{},
		&impl.GPIOPin{Pin: rpio.Pin(25)},
		&impl.GPIOPin{Pin: rpio.Pin(27)},
		&impl.GPIOPin{Pin: rpio.Pin(24)},
		ST7789.Screen240X240,
	)

	canvas := device.GetFullScreenCanvas()

	draw := gg.NewContext(240, 240)
	draw.SetRGB(255, 0, 0)
	draw.DrawRectangle(0, 0, 240, 240)
	draw.Fill()

	canvas.DrawImage(draw.Image())
	canvas.Flush()
}
