// Channels 是一種有類型的水管，透過 channel，你可以用 <- 來傳送與接收資料。
// ch <- v    // 傳送資料 v 到 channel ch
// v := <-ch  // 從 ch 接收資料並且指派給 v
// (其實資料就是跟著箭頭的方向流動)
//
// 就像 maps 和 slices 一樣，channels 必須在使用前先建立
// ch := make(chan int)
//
// 一般來說，直到另外一端準備好了，才會傳送和接收資料，不然就會先堵塞住 (block)
// 這讓 goroutines 不用顯式的 lock 或其他條件判斷就能做到同步
//
// 補充說明：可以把 <- 跟 <- 想成是一種 block 運算子
//          當程式執行到 <-c 時，它會停下來等待，直到有東西被送過來，這邊能接收為止
//          當程式執行到 c<-v 時，它也會停下來等，直到接收端可以接收，它才會往下繼續執行
//          藉此，就自然地達成了“同步”，而不需要其他的判斷
//
// The Go Approach:
// Don't communicate by sharing memory, share memory by communicating.
// from http://talks.golang.org/2012/concurrency.slide

package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
