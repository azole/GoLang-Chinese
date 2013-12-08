package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// struct 的欄位可以用.來存取
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // {4 2}
}
