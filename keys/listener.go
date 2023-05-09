package keys

import (
	"fmt"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"time"
)

func Listen(name string) {
	//if _, err := host.Init(); err != nil {
	//	panic(err)
	//}

	p := gpioreg.ByName(name)
	if p == nil {
		panic("failed to register pin")
	}

	if err := p.In(gpio.PullUp, gpio.RisingEdge); err != nil {
		panic(err)
	}

	t := time.Now()

	for {
		p.WaitForEdge(-1)
		t2 := time.Now()
		d := t2.Sub(t)
		if d.Milliseconds() < 170 {
			continue
		}
		fmt.Printf("%vms\n", d.Milliseconds())
		t = t2
	}
}
