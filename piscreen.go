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

	dc, rst, bl := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)
	dc.Output()
	rst.Output()
	bl.Output()
	bl.High()

	rst.High()
	time.Sleep(10 * time.Millisecond)
	rst.Low()
	time.Sleep(10 * time.Millisecond)
	rst.High()
	time.Sleep(10 * time.Millisecond)
}
