/* Go 的基本形態有：
 *        bool, string
 *        int, int8, int16, int32, int64
 *        uint, uint8, uint16, uint32, uint64, uintptr
 *        byte  // uint8 的別名
 *        rune  // int32 的別名
 *        float32, float64
 *        complex64, complex128  --> 令人驚奇的有複數型別
 */
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T (%v)\n"     // bool (false)
	fmt.Printf(f, ToBe, ToBe) // unit64 (18446744073709551615)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z) // complex128 ((2+3i))
}

// %T 是指輸出型別
// %v 是指輸出值
