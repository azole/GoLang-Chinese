// map 是一種 key/value 的資料結構
// map 在使用前必須由 make (而不是 new ) 來建立
// nil map 是空的，而且不能被指派

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	        	      // ^^^^ 一定要有逗點
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
}
