package impl

import "github.com/stianeikeland/go-rpio/v4"

type SPI struct{}

func (spi *SPI) SpiSpeed(speed uint32) {
	rpio.SpiSpeed(int(speed))
}

func (spi *SPI) SetSpiMode3() {
	rpio.SpiMode(1, 1)
}

func (spi *SPI) SpiTransmit(data []byte) {
	rpio.SpiTransmit(data...)
}
