package spi

import (
	"github.com/stianeikeland/go-rpio/v4"
	"image"
	"image/color"
	"time"
)

type ST7789 struct {
	width, height int

	dcPin, rstPin, blPin *rpio.Pin
}

func NewST7789() ST7789 {
	dcPin, rstPin, blPin := rpio.Pin(25), rpio.Pin(27), rpio.Pin(24)
	inst := ST7789{
		width:  240,
		height: 240,
		dcPin:  &dcPin,
		rstPin: &rstPin,
		blPin:  &blPin,
	}

	if err := rpio.Open(); err != nil {
		panic(err)
	}

	dcPin.Output()
	rstPin.Output()
	blPin.Output()
	blPin.High()

	inst.Reset()

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
	rpio.SpiChipSelect(0)
	rpio.SpiSpeed(40000000)

	inst.Command(SLPOUT)
	time.Sleep(1200 * time.Millisecond)

	inst.Command(MADCTL)
	inst.Data(0x70)

	inst.Command(COLMOD)
	inst.Data(0x05)

	inst.Command(PORCTRL)
	inst.Data(0x0C)
	inst.Data(0x0C)
	inst.Data(0x00)
	inst.Data(0x33)
	inst.Data(0x33)

	inst.Command(GCTRL)
	inst.Data(0x35)

	inst.Command(VCOMS)
	inst.Data(0x37)

	inst.Command(LCMCTRL)
	inst.Data(0x2C)

	inst.Command(VDVVRHEN)
	inst.Data(0x01)

	inst.Command(VRHS)
	inst.Data(0x12)

	inst.Command(VDVS)
	inst.Data(0x20)

	inst.Command(FRCTRL2)
	inst.Data(0x0F)

	inst.Command(PWCTRL1)
	inst.Data(0xA4)
	inst.Data(0xA1)

	inst.Command(PVGAMCTRL)
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

	inst.Command(NVGAMCTRL)
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

	inst.Command(INVON)

	inst.Command(DISPON)

	return inst
}

func (d ST7789) Command(cmds ...byte) {
	d.dcPin.Low()
	rpio.SpiTransmit(cmds...)
}

func (d ST7789) Data(data ...byte) {
	d.dcPin.High()
	rpio.SpiTransmit(data...)
}

func (d ST7789) Reset() {
	d.rstPin.High()
	time.Sleep(10 * time.Millisecond)
	d.rstPin.Low()
	time.Sleep(10 * time.Millisecond)
	d.rstPin.High()
	time.Sleep(10 * time.Millisecond)
}

func (d ST7789) SetWindows(x0, y0, x1, y1 byte) {
	d.Command(CASET)
	d.Data(0x00)
	d.Data(x0 & 0xFF)
	d.Data(0x00)
	d.Data((x1 - 1) & 0xFF)

	d.Command(RASET)
	d.Data(0x00)
	d.Data(y0 & 0xFF)
	d.Data(0x00)
	d.Data((y1 - 1) & 0xFF)

	d.Command(RAMWR)
}

func (d ST7789) Close() {
	rpio.SpiEnd(rpio.Spi0)
	rpio.Close()
}

func (d ST7789) Clear() {
	d.SetWindows(0, 0, 240, 240)
	c0 := RGBATo565(color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF})
	c1 := byte(c0)
	c2 := byte(c0 >> 8)
	for i := 0; i < 240*240; i++ {
		d.Data(c1, c2)
	}
}

func (d ST7789) ShowImage(img image.Image) {
	d.SetWindows(0, 0, 240, 240)
	for x := 0; x < 240; x++ {
		for y := 0; y < 240; y++ {
			c := RGBATo565(img.At(x, y))
			d.Data(byte(c), byte(c>>8))
		}
	}
}

func RGBATo565(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF8) + ((g & 0xFC) >> 5) + ((b & 0xF8) >> 11))
}
