package main

import (
	"fmt"
	"image/color"
)

func main() {
	colors := []color.RGBA{
		{R: 0xFF},
		{G: 0xFF},
		{B: 0xFF},
		{R: 0xFF, G: 0xFF, B: 0xFF},
		{R: 0xFF, B: 0xFF},
		{G: 0xFF, B: 0xFF},
		{R: 0x35, G: 0x37, B: 0x39},
	}

	for _, c := range colors {
		fmt.Printf("%+v\n", c)
		b := method1(c)
		fmt.Printf("  Method 1 = %X %X (0x%X)\n\n", byte(b>>8), byte(b), b)
		b = method2(c)
		fmt.Printf("  Method 2 = %X %X (0x%X)\n\n", byte(b>>8), byte(b), b)
		//b = method3(c)
		//fmt.Printf("  Method 3 = %X %X (0x%X)\n\n", byte(b), byte(b>>8), b)
	}

	r := 0x00
	g := 0xFF
	b := 0x00

	f1 := (r & 0b11111000) | (g >> 5)
	f2 := (g & 0b00011100) | (b & 0b11111000)
	f3 := (f1<<8 | f2) & 0x0FFF
	fmt.Printf("%X %X %X\n", f1, f2, f3)

	//c := colors[1]
	//r, g, b, _ := c.RGBA()
	//r = r & 0b11111000
	//fmt.Printf("%v %X %b\n", r, r, r)
	//g = g & 0b11111100
	//fmt.Printf("%v %X %b\n", g, g, g)
	//b = b & 0b11111000
	//fmt.Printf("%v %X %b\n", b, b, b)
	//
	//println()
	//
	//fmt.Printf("%v %X %b\n", r<<8, r<<8, r<<8)
	//fmt.Printf("%v %X %b\n", g<<3, g<<3, g<<3)
	//fmt.Printf("%v %X %b\n", b>>3, b>>3, b>>3)
	//fmt.Printf("%v %X %b\n", r<<8|g<<3, r<<8|g<<3, r<<8|g<<3)
	//
	//println()
	//
	//fmt.Printf("%v %X %b\n", r, r, r)
	//fmt.Printf("%v %X %b\n", g>>5, g>>5, g>>5)
	//fmt.Printf("%v %X %b\n", b>>11, b>>11, b>>11)
	//fmt.Printf("%v %X %b\n", r|g>>5, r|g>>5, r|g>>5)
	//
	//println()
	//
	//b = 0xE0
	//fmt.Printf("%v %X %b\n", b, b, b)
	//b = 0x07
	//fmt.Printf("%v %X %b\n", b, b, b)
	//b = 0xE0 | 0x07
	//fmt.Printf("%v %X %b\n", b, b, b)
}

func method1(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16(((r / 0xFF * 31) << 11) | ((g / 0xFF * 63) << 5) | (b / 0xFF * 31))
}

func method2(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16(((r & 0b11111000) << 8) | ((g & 0b11111100) << 3) | ((b & 0b11111000) >> 3))
}

//func method3(c color.RGBA) uint16 {
//	r, g, b, _ := c.RGBA()
//	return uint16((((r >> 3) & 0x1F) << 11) | (((g >> 2) & 0x3F) << 5) | ((b >> 3) & 0x1F))
//}
