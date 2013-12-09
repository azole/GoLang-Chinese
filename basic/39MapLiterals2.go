package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 如果最上層的類型只有類型名的話，元素中的實字就可以省略類型名
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967}, // 省略了 Vertex
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
