// go 並沒有 class，但可以替 struct 定義方法 (Type Methods)
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

// 補充說明
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.5.md
//
// "A method is a function with an implicit first argument, called a receiver." -- Rob Pike
// method 的名字一樣，但 receiver 不一樣時，那就是不同的 method
//
// 補充說明2
// 在範例 25Z_AnonymousFieldsInStruct.go 中有提到，
// 可以藉由匿名字段來繼承/組合，方法這邊一樣也可以，
// 當匿名字段是一個struct，他的方法可以被繼承下來。
// 當有衝突的時候，一樣也是外層優先
//
