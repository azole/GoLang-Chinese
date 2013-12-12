// 事實上，你可以在自己的 package 幫任何類型定義方法，而不僅止于 struct
// 但你不能幫其他 package 的類型定義方法

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
