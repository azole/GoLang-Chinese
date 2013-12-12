package main

import (
	"fmt"
)

// 請在這邊實作牛頓法的開方函式
// 牛頓法是先選擇一個起點 z，然後通過這個起點 z，根據以下公式反覆重複計算 z
// z = z - (z^2 -x ) / 2*z
// (1) 重複十次，觀察每次的產出，看答案是怎麼逼近的
// (2) 把條件改為當 z 沒有變化時停止，並且記錄這樣迭代了幾次，是否有比 10 次還少呢？
// 最後，還可以與 math.Sqrt 進行一下比較
// Hint: 要宣告一個浮點數可以 z:=float64(1) 或是 z:=1.0

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(i, z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	diff, z, i := 0.1, 1.0, 1
	for diff > 0.000000000000001 { // 這個差距是從 Sprt1 中觀察出來的
		i++
		diff = z
		z = z - (z*z-x)/(2*z)
		diff = diff - z
		if diff < 0 {
			diff = -diff
		}
	}
	fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt2(2))
}

// math.Sqrt(2) = 1.4142135623730951
