package main

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/stianeikeland/go-rpio/v4"
	"golang.org/x/image/bmp"
	"time"
)

var (
	dcPin, rstPin, blPin = rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	dcPin.Output()
	rstPin.Output()
	blPin.Output()
	blPin.High()

	rstPin.High()
	time.Sleep(10 * time.Millisecond)
	rstPin.Low()
	time.Sleep(10 * time.Millisecond)
	rstPin.High()
	time.Sleep(10 * time.Millisecond)

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)

	dcPin.Low()
	rpio.SpiTransmit(0x11)              // exit sleep
	time.Sleep(1200 * time.Millisecond) // allow to wake up

	rpio.SpiTransmit(0x36) // direction of frame memory read/write
	dcPin.High()
	rpio.SpiTransmit(0x70)

	dcPin.Low()
	rpio.SpiTransmit(0x3A)
	dcPin.High()
	rpio.SpiTransmit(0x05)

	dcPin.Low()
	rpio.SpiTransmit(0xB2)
	dcPin.High()
	rpio.SpiTransmit(0x0C)
	rpio.SpiTransmit(0x0C)
	rpio.SpiTransmit(0x00)
	rpio.SpiTransmit(0x33)
	rpio.SpiTransmit(0x33)

	dcPin.Low()
	rpio.SpiTransmit(0xB7)
	dcPin.High()
	rpio.SpiTransmit(0x35)

	dcPin.Low()
	rpio.SpiTransmit(0xBB)
	dcPin.High()
	rpio.SpiTransmit(0x37)

	dcPin.Low()
	rpio.SpiTransmit(0xC0)
	dcPin.High()
	rpio.SpiTransmit(0x2C)

	dcPin.Low()
	rpio.SpiTransmit(0xC2)
	dcPin.High()
	rpio.SpiTransmit(0x01)

	dcPin.Low()
	rpio.SpiTransmit(0xC3)
	dcPin.High()
	rpio.SpiTransmit(0x12)

	dcPin.Low()
	rpio.SpiTransmit(0xC4)
	dcPin.High()
	rpio.SpiTransmit(0x20)

	dcPin.Low()
	rpio.SpiTransmit(0xC6)
	dcPin.High()
	rpio.SpiTransmit(0x0F)

	dcPin.Low()
	rpio.SpiTransmit(0xD0)
	dcPin.High()
	rpio.SpiTransmit(0xA4)
	rpio.SpiTransmit(0xA1)

	dcPin.Low()
	rpio.SpiTransmit(0xE0)
	dcPin.High()
	rpio.SpiTransmit(0xD0)
	rpio.SpiTransmit(0x04)
	rpio.SpiTransmit(0x0D)
	rpio.SpiTransmit(0x11)
	rpio.SpiTransmit(0x13)
	rpio.SpiTransmit(0x2B)
	rpio.SpiTransmit(0x3F)
	rpio.SpiTransmit(0x54)
	rpio.SpiTransmit(0x4C)
	rpio.SpiTransmit(0x18)
	rpio.SpiTransmit(0x0D)
	rpio.SpiTransmit(0x0B)
	rpio.SpiTransmit(0x1F)
	rpio.SpiTransmit(0x23)

	dcPin.Low()
	rpio.SpiTransmit(0xE1)
	dcPin.High()
	rpio.SpiTransmit(0xD0)
	rpio.SpiTransmit(0x04)
	rpio.SpiTransmit(0x0C)
	rpio.SpiTransmit(0x11)
	rpio.SpiTransmit(0x13)
	rpio.SpiTransmit(0x2C)
	rpio.SpiTransmit(0x3F)
	rpio.SpiTransmit(0x44)
	rpio.SpiTransmit(0x51)
	rpio.SpiTransmit(0x2F)
	rpio.SpiTransmit(0x1F)
	rpio.SpiTransmit(0x1F)
	rpio.SpiTransmit(0x20)
	rpio.SpiTransmit(0x23)

	dcPin.Low()
	rpio.SpiTransmit(0x21)

	rpio.SpiTransmit(0x29)

	rpio.SpiTransmit(0x2A)
	dcPin.High()
	rpio.SpiTransmit(0x00)
	rpio.SpiTransmit(0 & 0xFF)
	rpio.SpiTransmit(0x00)
	rpio.SpiTransmit((240 - 1) & 0xFF)

	dcPin.Low()
	rpio.SpiTransmit(0x2B)
	dcPin.High()
	rpio.SpiTransmit(0x00)
	rpio.SpiTransmit(0 & 0xFF)
	rpio.SpiTransmit(0x00)
	rpio.SpiTransmit((240 - 1) & 0xFF)

	dcPin.High()
	rpio.SpiTransmit(0x2C)

	context := gg.NewContext(240, 240)
	context.SetRGB(255, 0, 0)
	context.DrawRectangle(0, 0, 240, 240)
	context.Fill()

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, context.Image()); err != nil {
		panic(err)
	}

	dcPin.High()
	rpio.SpiTransmit(buf.Bytes()...)

	rpio.SpiEnd(rpio.Spi0)

	print(len(buf.Bytes()))
	print(len(buf.Bytes()) / 4096)
}
