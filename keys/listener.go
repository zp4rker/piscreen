package keys

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func Listen(name string) {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	p := gpioreg.ByName(name)
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
