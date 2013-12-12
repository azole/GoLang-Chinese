// 練習題：實作 Pic 這個函式
// 這個函式會回傳一個長度為 dy 的 slice
// 這個 slice 中每個元素是一個 dx 8-bit unsigned integers 的 slice
// 當執行這個程式的時候，他會展示一張被轉為灰度的圖片。
//
// 要選擇哪一張圖片都可以。一些有趣的轉換函式像是：x^y, (x+y)/2, 和 x*y
//
// 這邊需要迭代每一個 []uint8，並且分配一個 []uint8 給他
// 利用 uint8(intValue) 去轉型

package main

import "code.google.com/p/go-tour/pic"

// 我不確定這是不是最好的寫法...
// 這三種函式做出來的圖都很美，值得跑跑看
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
		for j := range p[i] {
			//p[i][j] = uint8(i^j)
			//p[i][j] = uint8((i+j)/2)
			p[i][j] = uint8(i * j)
		}
	}
	return p
}

func main() {
	pic.Show()
}
