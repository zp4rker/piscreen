package keys

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
	"piscreen/screens"
	"piscreen/vars"
	"time"
)

func StartKeyListeners() {
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
					t := time.Now()
					if t.Sub(k.LastRegistered).Milliseconds() < int64(vars.ListenDelay) {
						continue
					}
					k.LastRegistered = t
					screens.CurrentScreen.Handle(k.Name)
				}
			}
		}(key)
	}
}
