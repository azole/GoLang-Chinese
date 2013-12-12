package main

import "fmt"

// 函式可以回傳任何數量的回傳值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
