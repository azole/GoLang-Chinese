// 數值型的常數是高精確度的值。
// 沒有型別的常數會依靠上下文來判斷其型別。

package main

import "fmt"

// 分組宣告
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big) // 編譯時會出現 constant 1267650600228229401496703205376 overflows int
	fmt.Println(Small)          // 2
	fmt.Println(needInt(Small)) // 21
	//fmt.Println(needInt(Big))     // 編譯時會出現 constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
}

// 補充說明：
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.2.md
//
// 可以用 const 的分組宣告來做列舉 enum，利用 itoa 這個關鍵字，itoa 開始的值為 0，每叫用一次就加 1
//
// const(
//     x = iota  // x == 0
//     y = iota  // y == 1
//     z = iota  // z == 2
//     w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 w = iota，因此 w == 3。其实上面y和z可同样不用"= iota"
// )
//
// const v = iota // 每遇到一个 const 关键字，iota 就会重置，此时 v == 0
//
// const (
//   e, f, g = iota, iota, iota //e=0,f=0,g=0 iota 在同一行值相同
// )
//
