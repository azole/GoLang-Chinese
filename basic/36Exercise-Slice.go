// 練習題：實作 Pic 這個函式
// 這個函式會回傳一個長度為 dy 的 slice
// 這個 slice 中每個元素是一個 dx 8-bit unsigned integers 的 slice
// 當執行這個程式的時候，他會展示一張被轉為灰度的圖片。
//
// 要選擇哪一張圖片都可以
// The choice of image is up to you. Interesting functions include x^y, (x+y)/2, and x*y.
//
// 這邊需要迭代每一個 []uint8，並且分配一個 []uint8 給他
// 利用 uint8(intValue) 去轉型

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show()
}
