// 可以藉由指派 _ 來省略 index 或 value
// 如果只想要 index，可以把 ", value" 整個省略掉

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
