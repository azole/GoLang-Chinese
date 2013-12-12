package main

import "fmt"

// struct 是一堆屬性的集合
// 這邊 type 正如你所想的那樣
type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
