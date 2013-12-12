// struct 實字透過列出欄位的值來表示一個新分配的 struct
// 也可以透過指定欄位名稱來只列出欄位的部分集合(例如只列出 X)，這時候順序就沒關係
// & 這個符號是會建立一個指標，指向新建立的 struct

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // {1 2} p 的型別是 Vertex
	q = &Vertex{1, 2} // &{1 2}  q 的型別是 *Vertex
	r = Vertex{X: 1}  // {1 0}  隱含 Y 是 0
	s = Vertex{}      // {0 0}  X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
