package keys

import (
	"fmt"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
	"piscreen/util"
	"piscreen/vars"
	"time"
)

func Listen() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	for _, key := range Keys {
		p := gpioreg.ByName(key.Pin)
		if p == nil {
			panic("failed to register pin " + key.Pin)
		}

		if err := p.In(gpio.PullUp, gpio.RisingEdge); err != nil {
			panic(err)
		}

		go func(k Key) {
			for vars.Running {
				if p.WaitForEdge(-1) {
					util.Debug(fmt.Sprintf("received press of %v", k.Name))
					t := time.Now()
					if t.Sub(k.LastRegistered).Milliseconds() < int64(vars.ListenDelay) {
						continue
					}
					util.Debug(fmt.Sprintf("pressed after %vms", t.Sub(k.LastRegistered).Milliseconds()))
					k.LastRegistered = t
					vars.CurrentScreen.Handle(k.Name)
					util.Debug("press handled")
				}
			}
		}(key)
	}
}
