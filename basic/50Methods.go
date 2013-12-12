// go 並沒有 class，但可以替 struct 定義方法
// method receiver 是寫在 func 跟 方法名稱的中間
// --> 這個 method receiver 有點像是 this 的感覺

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
