// 省略前後兩個分號後，看起來有沒有很像 C 中的 while
// 所以，雖然 go 中只有 for 這種回圈結構，但其實也能做到像 while 一樣的事情
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
