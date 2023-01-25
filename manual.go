package main

import (
	"github.com/stianeikeland/go-rpio/v4"
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

	cmd(0x11)                           // exit sleep
	time.Sleep(1200 * time.Millisecond) // allow to wake up

	cmd(0x36) // direction of frame memory read/write
	data(0x70)

	cmd(0x3A)
	data(0x05)

	cmd(0xB2)
	data(0x0C)
	data(0x0C)
	data(0x00)
	data(0x33)
	data(0x33)

	cmd(0xB7)
	data(0x35)

	cmd(0xBB)
	data(0x37)

	cmd(0xC0)
	data(0x2C)

	cmd(0xC2)
	data(0x01)

	cmd(0xC3)
	data(0x12)

	cmd(0xC4)
	data(0x20)

	cmd(0xC6)
	data(0x0F)

	cmd(0xD0)
	data(0xA4)
	data(0xA1)

	cmd(0xE0)
	data(0xD0)
	data(0x04)
	data(0x0D)
	data(0x11)
	data(0x13)
	data(0x2B)
	data(0x3F)
	data(0x54)
	data(0x4C)
	data(0x18)
	data(0x0D)
	data(0x0B)
	data(0x1F)
	data(0x23)

	cmd(0xE1)
	data(0xD0)
	data(0x04)
	data(0x0C)
	data(0x11)
	data(0x13)
	data(0x2C)
	data(0x3F)
	data(0x44)
	data(0x51)
	data(0x2F)
	data(0x1F)
	data(0x1F)
	data(0x20)
	data(0x23)

	cmd(0x21)

	cmd(0x29)
}

func cmd(cmds ...byte) {
	dcPin.Low()

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)

	rpio.SpiTransmit(cmds...)

	rpio.SpiEnd(rpio.Spi0)
}

func data(data ...byte) {
	dcPin.High()

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)

	rpio.SpiTransmit(data...)

	rpio.SpiEnd(rpio.Spi0)
}
