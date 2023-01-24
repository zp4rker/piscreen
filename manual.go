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
