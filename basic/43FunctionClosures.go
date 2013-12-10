// 這不就是傳說中的 javascript closure 嗎？
// 有興趣的朋友歡迎參考：http://imazole.wordpress.com/tag/closure/
//
// go 函式也可以是 closures
// closure 是一個函式值，它會"參考"到在函式本體之外的變數
// 函式可以存取跟指派這個參考的變數
// 通常會說，這個函式被綁定在這些變數上
//
// 例如，adder 就會回傳一個 closure
// 每一個 closure 會綁定他自己的 sum 變數

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
