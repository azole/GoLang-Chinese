// go 只有 for 這種迴圈結構，而且很像 C 或是 Java 的 for，只是沒有了 ()，但 {} 還是需要的。

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
