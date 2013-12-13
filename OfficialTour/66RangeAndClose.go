// 傳送者可以 close 一個 channel，用來表示不會再有資料被送出了
// 接收者可以藉由接收表示式的第二個參數來知道一個 channel 是否已經被關閉了
// v, ok := <-ch
// 當 ok 為 false 時，表示已經沒有任何資料可以接收了，這個 channel 已經被關閉了
// The loop for i := range c receives values from the channel repeatedly until it is closed.
// 迴圈中的 for i:= range c 會反覆地從 channel 接收資料，直到它被關閉了
//
// note1: 只有傳送者可以關閉一個 channel，接收者永遠不行。
// 如果要對一個已經關閉的 channel 傳送資料會發生 panic.
//
// note2: channels 不像檔案，通常你不需要特別關閉它
// 只有在要告訴接收者再也沒有資料會送過來時，才需要去關閉它
// 例如下面的例子，告訴 range loop 終止了

package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	// 如果不加 close(c) 執行時會發生錯誤
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:

}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

}
