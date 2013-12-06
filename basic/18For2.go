// 跟 C 或 Java 一樣，可以讓 for 的第一個跟最後一個敘述為空
// 可以寫成 for ; sum < 1000; {
// 但 goSublime太聰明了，一直幫我把前後兩個 ; 消掉
// 變成像範例 19 那樣的 while 的形式

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
