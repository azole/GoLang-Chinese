// go 有 pointers，但沒有指標的運算
// struct 的欄位可以透過 struct 指標來存取
// The indirection through the pointer is transparent. -> 這句不知道什麼意思

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}
