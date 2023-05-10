package core

import (
	"image"
	"image/color"
)

type Display interface {
	Command(cmds ...byte)
	Data(data ...byte)

	Close()
	Reset()
	SetWindows(x0, y0, x1, y1 byte)

	Clear(c color.RGBA)
	ToggleSleep()
	IsAsleep() bool
	ShowImage(img image.Image)
}
