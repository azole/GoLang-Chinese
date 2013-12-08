// 跟 C 或 Java 一樣，可以讓 for 的第一個跟最後一個敘述為空
// 可以寫成 for ; sum < 1000; {

package main

import "fmt"

func main() {
	sum := 1
	for ;sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
