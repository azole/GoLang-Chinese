// image package 定義了一些 Image 的介面
// package image
//
// type Image interface {
//   ColorModel() color.Model
//   Bounds() Rectangle
//   At(x, y int) color.Color
// }
//
// 同樣的，color.Color 和 color.Model 也都是介面，
// 但我們通常忽略它們，因為我們都常都使用預先實作好的 color.RGBA 跟 color.RGBAModel

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
