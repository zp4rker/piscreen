package main

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/rubiojr/go-pirateaudio/display"
	"golang.org/x/image/bmp"
	"image/color"
	"time"
)

func main() {
	dsp, err := display.Init()
	if err != nil {
		panic(err)
	}
	defer dsp.Close()

	context := gg.NewContext(240, 240)
	context.SetRGB(255, 0, 0)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, context.Image()); err != nil {
		panic(err)
	}

	// Set the screen color to white
	dsp.FillScreen(color.RGBA{R: 0, G: 0, B: 0, A: 0})

	// Rotate before pushing pixels, so the image appears rotated
	dsp.Rotate(display.ROTATION_180)
	dsp.DrawImage(buf)

	println("Anything happen?")
	time.Sleep(5 * time.Second)
	println("How about now?")
}
