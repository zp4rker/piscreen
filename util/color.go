package util

import "github.com/fogleman/gg"

func GGColor(r, g, b int) (float64, float64, float64) {
	red := float64(r) / 0xFF
	green := float64(g) / 0xFF
	blue := float64(b) / 0xFF
	return red, green, blue
}

func White(context *gg.Context) {
	context.SetRGB(GGColor(0xFF, 0xFF, 0xFF))
}

func Black(context *gg.Context) {
	context.SetRGB(GGColor(0x00, 0x00, 0x00))
}

func Grey(context *gg.Context) {
	context.SetRGB(GGColor(0xC2, 0xC2, 0xC2))
}

func DarkGrey(context *gg.Context) {
	context.SetRGB(GGColor(0x35, 0x37, 0x39))
}
