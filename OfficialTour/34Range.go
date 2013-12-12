// for 迴圈的 range 格式可以迭代 slice 或 map
// 語法是 ( ExpressionList "=" | IdentifierList ":=" ) "range" Expression
// range 的右邊的這個 Expression 可以是 array, slice, map, string, channel 等等
// 如果是 slice 或 array 的時候，左側的兩個值會是 index 跟 value
// 如下範例所示：

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
*/
