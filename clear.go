package main

import (
	"piscreen/spi"
)

func main() {
	disp := spi.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear()
	println("done")
}
