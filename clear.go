package main

import (
	"image/color"
	"piscreen/spi"
	"piscreen/util"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	util.Debug("clearing...\n")
	disp.Clear(color.RGBA{})
	util.Debug("done\n")
}
