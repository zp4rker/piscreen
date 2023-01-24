package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(13)
	pin.Input()
	pin.PullUp()
	fmt.Printf("PullUp: %d, %d\n", pin.Read(), pin.ReadPull())
}
