package keys

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func Listen() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	p := gpioreg.ByName("GPIO13")
	if p == nil {
		panic("failed to register pin")
	}

	if err := p.In(gpio.PullUp, gpio.FallingEdge); err != nil {
		panic(err)
	}

	for {
		p.WaitForEdge(-1)
		println("Got press!")
	}
}
