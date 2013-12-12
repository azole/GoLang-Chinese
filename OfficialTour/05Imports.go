package main

// go 如果要 import 多個 package，可以用括號來一次引入多個
import (
	"fmt"
	"math"
)

// 或是這樣寫也是一樣的效果
// import "fmt"
// import "math"

func main() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}
