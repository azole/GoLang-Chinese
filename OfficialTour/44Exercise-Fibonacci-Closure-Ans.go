package main

import "fmt"

// 實作一個計算 fibonacci 的函式

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
	a, b := 0, -1
	return func() int {
		if b == -1 {
			b++
			return b
		} else if b == 0 {
			b++
			return b
		} else {
			c := b
			b = a + b
			a = c
			return b
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
