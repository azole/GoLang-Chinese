// go 有 pointers，但沒有指標的運算
// struct 的欄位可以透過 struct 指標來存取
// The indirection through the pointer is transparent. -> 這句不知道什麼意思
// 我猜是指，透過 pointer 的存取，會影響到原本的值
// 例如以下範例，修改了 q，但 p 也會被影響
// 而 z 不是透過 pointer 存取，所以 z 的修改不會影響到 p

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	z := p
	q.X = 1e9
	z.Y = 55
	fmt.Println(p) // {1000000000 2}
	fmt.Println(z) // {1 55}
}
