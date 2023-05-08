package main

import (
	"piscreen/spi/impl"
)

func main() {
	disp := impl.NewST7789()
	defer disp.Close()

	println("clearing...")
	disp.Clear()
	println("done")
}
