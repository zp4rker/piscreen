package main

import "github.com/stianeikeland/go-rpio/v4"

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	dcPin, rstPin, blPin := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)

	defer rpio.Close()
}
