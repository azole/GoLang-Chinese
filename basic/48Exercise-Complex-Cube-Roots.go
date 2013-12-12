// 通過 complex64, complex128 來探索一下 go 內建的複數型態
// 為了找到立方根，牛頓法需要大量的循環
// z = z - ((z^3-x)/(3*z^2))
//
// 提示：math/cmplx 裡有 Pow 函式

package main

import "fmt"

func Cbrt(x complex128) complex128 {
}

func main() {
	fmt.Println(Cbrt(2))
}
