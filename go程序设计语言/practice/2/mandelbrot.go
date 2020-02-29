package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	pWidth, pHeight        = 4096 * 2, 4096 * 2
	avgSize                = 2
	RGBACharSize           = 3
)

func main() {
	fptr, _ := os.Create("test.png")

	image := image.NewRGBA(image.Rect(0, 0, pWidth, pHeight))
	for py := 0; py < pHeight; py += avgSize {
		for px := 0; px < pWidth; px += avgSize {
			findAverageColor(image, px, py)
		}
	}
	png.Encode(fptr, image)
}	  

func findAverageColor(image *image.RGBA, px, py int) {
	s := [][]uint8{}
	for m := 0; m < avgSize; m++ {
		y := float64(py+m)/pHeight*(ymax-ymin) + ymin
		char := []uint8{}
		for n := 0; n < avgSize; n++ {
			x := float64(px+n)/pWidth*(xmax-xmin) + xmin
			z := complex(x, y)
			char = mandelbrot(z)
			s = append(s, char)
		}
	}

	size := avgSize * avgSize
	tmp := make([]int, RGBACharSize)
	for i := 0; i < RGBACharSize; i++ {
		for j := 0; j < size; j++ {
			tmp[i] += int(s[j][i])
		}
	}

	for m := 0; m < avgSize; m++ {
		for n := 0; n < size; n++ {
			image.Set(px+n, py+m, color.RGBA{
				uint8(tmp[0] / size),
				uint8(tmp[1] / size),
				uint8(tmp[1] / size),
				128,
			})
		}
	}
}

func mandelbrot(z complex128) []uint8 {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return []uint8{uint8(0 + contrast*n), uint8(0 + contrast*n), uint8(0 + contrast*n)}
		}
	}

	return []uint8{0, 0, 0}
}
