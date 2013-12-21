// 補充說明 (官方 Tour 中沒有)
// 來源：http://blog.golang.org/defer-panic-and-recover
//
// go 中沒有一般的異常處理，它是利用 defer / panic / recover 來做
// panic 像是我們其他語言中丟出異常一樣
// 然後在 defer 中用 recover 來接

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// 如果這邊不用 panic 來接 panic 的話，程式就會中斷
	// recover() 回傳的是 panic 傳出來的東西
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("[test %v]", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
