// 通過 complex64, complex128 來探索一下 go 內建的複數型態
// 為了找到立方根，牛頓法需要大量的循環
// z = z - ((z^3-x)/(3*z^2))
//
// 提示：math/cmplx 裡有 Pow 函式

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var z, z1 complex128 = 1.0, 1.0
	diff := 1.0
	i := 0
	for diff > 0.0 {
		i++
		z1 = z
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
		im := imag(z) - imag(z1)
		rl := real(z) - real(z1)
		diff = math.Sqrt(im*im + rl*rl)
		//fmt.Println(i, z, diff)
	}
	return z
}

func main() {
	ans := Cbrt(2)
	fmt.Println(ans)
	fmt.Println(cmplx.Pow(ans, 3))
}
