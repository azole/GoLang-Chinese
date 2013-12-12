// 跟 for 一樣，在執行條件判斷之前，可以放一些短敘述。
// 由這個敘述宣告的變數，其範疇只有在這個 if 範圍內。

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// 這邊如果要用 v，編譯時會出現 undefined v
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20)) // 9 20
}
