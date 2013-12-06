// 在短敘述內宣告出來的變數，在 esle 範圍內一樣可以被使用

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 這邊如果要用 v，編譯時會出現 undefined v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // 這邊結尾也要一個 ,
	)
}

// 印出來的順序有點奇特
// 27 >= 20
// 9 20
// 所以 Println 會等參數中的函式呼叫全部都執行完，才開始列印？
