package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	W, H int
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.W, im.H)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
