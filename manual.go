package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	dcPin, rstPin, blPin := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)

	dcPin.Output()
	rstPin.Output()
	blPin.Output()
	blPin.High()

	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)

	rstPin.High()
	time.Sleep(10 * time.Millisecond)
	rstPin.Low()
	time.Sleep(10 * time.Millisecond)
	rstPin.High()
	time.Sleep(10 * time.Millisecond)
}
