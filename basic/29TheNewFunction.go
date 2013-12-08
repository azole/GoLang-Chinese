// 透過 new 函式可以建立一個零初始化值的 T 值，並回傳其指標
// 語法為：
// var t *T = new(T)
// t := new(T)

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v) // &{0 0}
	v.X, v.Y = 11, 9
	fmt.Println(v) // &{11 9}
}
