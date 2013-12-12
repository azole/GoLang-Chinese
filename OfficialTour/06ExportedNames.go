package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}

// 編譯後會出現錯誤：cannot refer to unexported name math.pi
// 在 Go 中，只有首字母大寫的才會被 export
// 這邊改成 math.Pi 就可以了
