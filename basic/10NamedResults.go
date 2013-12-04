package main

import "fmt"

// 在 Go 中，函式可以有多個回傳值，而且還能被命名
// 如果回傳值有被命名，則一個沒有參數的 return，
// 會回傳這些回傳值當時的值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 執行到這行時，x, y 分別為多少就會回傳其值
}

func main() {
	fmt.Println(split(23))
}
