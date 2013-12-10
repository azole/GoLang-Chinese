package main

import "fmt"

// 實作一個計算 fibonacci 的函式

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
