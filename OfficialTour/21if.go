// 跟 for 迴圈一樣，go 的 if 也很像 C 跟 Java 的 if，只是少了 ()，但 {} 仍是需要的
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i
}
