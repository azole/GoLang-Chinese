// Channels 是可以被緩衝的。make 的第二個參數就是 buffer 的長度，可以用來初始化可緩衝的 channel
// ch := make(chan int, 100)
//
// 只有當 buffer 是滿的時候，傳送到該 channel 的部分會被堵塞住
// 只有當 buffer 是空的時候，接收該 channel 才會被堵塞住
// --> 也就是 buffered channels 把同步的特性移掉了，只有在滿或空的時候才會出現堵塞 (block)
//
// 修改一下範例，看看超過 buffer 時會怎麼樣
//

package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3
	// 如果多加這一行，執行時會發生錯誤
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:

	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c)
	// 如果多加這一行，執行時會發生錯誤
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
}
