// 介面是一組方法的定義
// 一個介面類型的值會存放任何想要被實作的方法

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	//a = v  // a Vertex, does NOT implements Abser
	// 這行會有錯誤發生

	fmt.Println(a.Abs())

	// 以下為補充說明：
	//value, ok := a.(MyFloat) // ok = false，因為 a 最後是存了 &v
	value, ok := a.(*Vertex) // ok = true
	//value, ok := a.(Vertex)	// 會有 Error，因為 Vertex 並沒有實作 Abser
	fmt.Println(value, ok)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 補充說明
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.6.md
//
// 空的 interface
// var a interface{}
// 不包含任何 method，所以沒有任何作用
// 但可以用來儲存任意型別的數值，有點像 C 的 void*
// 例如，一個函式把 interface{} 當作參數，那這個函式就可以接收任意型態的參數
// 同理，如果一個函式回傳 interface{}，那這個函式就能回傳任意型態的值
//
// 補充說明
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.6.md
//
// 判斷 interface 變量儲存的類型
// 以上面的範例來說，a 是 Abser 這個介面類型，它可以儲存有實作該介面的 MyFloat 跟 *Vertex
// 可以用 a.(MyFloat) 跟 a.(*Vertex) 來判斷目前 a 是儲存哪一種型態的值
//
// 如果可以儲存的型態有很多種 (也就是有很多種 struct 有實作該介面)，
// 判斷時難免用上一堆 if-esle，這時候可以用 switch 來取代
// switch value := element.(type) {}
// 注意：element.(type) 不能用在 switch 以外的地方
