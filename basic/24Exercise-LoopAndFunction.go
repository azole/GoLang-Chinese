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

func Sqrt(x float64) float64 {
}

func main() {
	fmt.Println(Sqrt(2))
}
