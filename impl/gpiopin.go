package impl

import "github.com/stianeikeland/go-rpio/v4"

type GPIOPin struct {
	rpio.Pin
}

func (p *GPIOPin) SetOutput() {
	p.Output()
}
