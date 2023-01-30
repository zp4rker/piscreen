package spi

import "image"

type Display interface {
	Command(cmds ...byte)
	Data(data ...byte)

	Reset()

	Clear()

	ShowImage(img image.Image)
}
