// 修改 Sqrt 函式，使其接受一個負數時，返回 ErrNegativeSqrt
//
// 首先，要先建立一個新類型：
// type ErrNegativeSqrt float64
// 然後幫它實現：
// func (e ErrNegativeSqrt) Error() string
//
// 注意：呼叫 fmt.Print(e) 時，會陷入無窮回圈
// 必須要這樣呼叫：fmt.Print(float64(e))
// 為什麼呢？

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
