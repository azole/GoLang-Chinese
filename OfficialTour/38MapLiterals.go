package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map 實字很像 struct，只是一定要有 key
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	//^^^ 一定要有逗點
}

func main() {
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
