package util

import "image"

func ImageCmp(img1, img2 image.Image) bool {
	if !img1.Bounds().Eq(img2.Bounds()) {
		return false
	}

	for x := 0; x < img1.Bounds().Size().X; x++ {
		for y := 0; y < img1.Bounds().Size().Y; y++ {
			c1 := img1.At(x, y)
			c2 := img2.At(x, y)
			r1, g1, b1, _ := c1.RGBA()
			r2, g2, b2, _ := c2.RGBA()
			if r1 != r2 {
				return false
			}
			if g1 != g2 {
				return false
			}
			if b1 != b2 {
				return false
			}
		}
	}

	return true
}
