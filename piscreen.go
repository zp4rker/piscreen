package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"piscreen/spi/impl"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	defer rpio.Close()

	impl.NewST7789()
}
