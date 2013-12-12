// method 可以跟命名類型關聯，或是指標命名類型關聯
// 前兩個 Abs 的例子中，我們看到其中一個是跟 *Vertex 指標類型，另外一個是 MyFloat 值類型。
// 這邊有兩個使用指標接收者的理由：
// 1. 避免每次方法呼叫時都做值的複製 (對大型 struct 來說，傳指標會比複製一份有效率)
// 2. 傳指標可以真的修改到原本的值
//
// 以下範例可以試著把指標拿掉，改用 Vertex取代
// 你就會發現 Scale 毫無作用
// 當 Scale 在處理 v 時，處理的是複製的一份，而不是原本的那份
// Abs 也是，只是 Abs 沒差，因為他不需要修改到原本的值

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
