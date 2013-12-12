package main

import "fmt"

// 函式可以有 0 或多個參數，要留意的是型別是放在變數的後面
// (這跟我們習慣的 C 系列語言不太一樣)
func add(x int, y int) int {
	return x + y
}

// 當兩個或連續多個參數是同一種型別時，
// 除了最後一個外，其他的都可以省略
func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(sub(7, 3))
}
