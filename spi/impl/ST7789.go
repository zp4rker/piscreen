package impl

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

type ST7789 struct {
	width, height int

	dcPin, rstPin, blPin *rpio.Pin
}

func (d ST7789) Command(cmds ...byte) {
	d.dcPin.Low()
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
	rpio.SpiTransmit(cmds...)
	rpio.SpiEnd(rpio.Spi0)
}

func (d ST7789) Data(data ...byte) {
	d.dcPin.High()
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
	rpio.SpiTransmit(data...)
	rpio.SpiEnd(rpio.Spi0)
}

func (d ST7789) Reset() {}

func NewST7789() {
	dcPin, rstPin, blPin := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)
	inst := ST7789{
		width:  240,
		height: 240,
		dcPin:  &dcPin,
		rstPin: &rstPin,
		blPin:  &blPin,
	}

	inst.Reset()

	inst.Command(0x11) // DISPON
	time.Sleep(1200 * time.Millisecond)
	inst.Command(0x36) // MADCTL
	inst.Data(0x70)

	inst.Command(0x3A) // COLMOD
	inst.Data(0x05)

	inst.Command(0xB2) // PORCTRL
	inst.Data(0x0C)
	inst.Data(0x0C)
	inst.Data(0x00)
	inst.Data(0x33)
	inst.Data(0x33)

	inst.Command(0xB7) // GCTRL
	inst.Data(0x35)

	inst.Command(0xBB) // VCOMS
	inst.Data(0x37)

	inst.Command(0xC0) // LCMCTRL
	inst.Data(0x2C)

	inst.Command(0xC2) // VDVVRHEN
	inst.Data(0x01)

	inst.Command(0xC3) // VRHS
	inst.Data(0x12)

	inst.Command(0xC4) // VDVS
	inst.Data(0x20)

	inst.Command(0xC6) // FRCTRL2
	inst.Data(0x0F)

	inst.Command(0xD0) // PWCTRL1
	inst.Data(0xA4)
	inst.Data(0xA1)

	inst.Command(0xE0) // PVGAMCTRL
	inst.Data(0xD0)
	inst.Data(0x04)
	inst.Data(0x0D)
	inst.Data(0x11)
	inst.Data(0x13)
	inst.Data(0x2B)
	inst.Data(0x3F)
	inst.Data(0x54)
	inst.Data(0x4C)
	inst.Data(0x18)
	inst.Data(0x0D)
	inst.Data(0x0B)
	inst.Data(0x1F)
	inst.Data(0x23)

	inst.Command(0xE1) // NVGAMCTRL
	inst.Data(0xD0)
	inst.Data(0x04)
	inst.Data(0x0C)
	inst.Data(0x11)
	inst.Data(0x13)
	inst.Data(0x2C)
	inst.Data(0x3F)
	inst.Data(0x44)
	inst.Data(0x51)
	inst.Data(0x2F)
	inst.Data(0x1F)
	inst.Data(0x1F)
	inst.Data(0x20)
	inst.Data(0x23)

	inst.Command(0x21) // INVON

	inst.Command(0x29) // DISPON
}
