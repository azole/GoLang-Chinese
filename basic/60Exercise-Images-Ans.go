package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	colr          uint8
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

func (m *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) At(x, y int) color.Color {
	return color.RGBA{m.colr + uint8(x), m.colr + uint8(y), 255, 255}
}

func main() {
	m := &Image{100, 100, 128}
	pic.ShowImage(m)
}
