package util

func GGColor(r, g, b int) (float64, float64, float64) {
	red := float64(r) / 0xFF
	green := float64(g) / 0xFF
	blue := float64(b) / 0xFF
	return red, green, blue
}
