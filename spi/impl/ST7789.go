package impl

import (
	"bytes"
	"github.com/stianeikeland/go-rpio/v4"
	"image"
	"image/png"
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
	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)
	rpio.SpiTransmit(cmds...)
	rpio.SpiEnd(rpio.Spi0)
}

func (d ST7789) Data(data ...byte) {
	d.dcPin.High()
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)
	rpio.SpiTransmit(data...)
	rpio.SpiEnd(rpio.Spi0)
}

func (d ST7789) Reset() {
	d.rstPin.High()
	time.Sleep(10 * time.Millisecond)
	d.rstPin.Low()
	time.Sleep(10 * time.Millisecond)
	d.rstPin.High()
	time.Sleep(10 * time.Millisecond)
}

func (d ST7789) ShowImage(img image.Image) {
	d.Command(0x2A) // CASET
	d.Data(0x00)
	d.Data(0 & 0xFF)
	d.Data(0x00)
	d.Data((240 - 1) & 0xFF)

	d.Command(0x2B) // RASET
	d.Data(0x00)
	d.Data(0 & 0xFF)
	d.Data(0x00)
	d.Data((240 - 1) & 0xFF)

	d.Command(0x2C) // RAMWR

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		panic(err)
	}

	d.Data(buf.Bytes()...)
}

func (d ST7789) Clear() {}

func NewST7789() ST7789 {
	dcPin, rstPin, blPin := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)
	inst := ST7789{
		width:  240,
		height: 240,
		dcPin:  &dcPin,
		rstPin: &rstPin,
		blPin:  &blPin,
	}

	dcPin.Output()
	rstPin.Output()
	blPin.Output()
	blPin.High()

	inst.Reset()

	inst.Command(0x11)
	time.Sleep(1200 * time.Millisecond)

	inst.Command(0x36)
	inst.Data(0x70)

	inst.Command(0x3A)
	inst.Data(0x05)

	inst.Command(0xB2)
	inst.Data(0x0C)
	inst.Data(0x0C)
	inst.Data(0x00)
	inst.Data(0x33)
	inst.Data(0x33)

	inst.Command(0xB7)
	inst.Data(0x35)

	inst.Command(0xBB)
	inst.Data(0x37)

	inst.Command(0xC0)
	inst.Data(0x2C)

	inst.Command(0xC2)
	inst.Data(0x01)

	inst.Command(0xC3)
	inst.Data(0x12)

	inst.Command(0xC4)
	inst.Data(0x20)

	inst.Command(0xC6)
	inst.Data(0x0F)

	inst.Command(0xD0)
	inst.Data(0xA4)
	inst.Data(0xA1)

	inst.Command(0xE0)
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

	inst.Command(0xE1)
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

	inst.Command(0x21)

	inst.Command(0x29)

	return inst
}
