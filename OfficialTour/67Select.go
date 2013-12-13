// select 可以讓 goroutine 在多個溝通運算上等待
// select 會堵塞，直到它的其中一個 case 可以被執行，然後它執行這個 case
// 如果同時有多個已經準備好，那它會隨機挑選一個

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() { // create a new goroutine
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit) // execute on current goroutine
}
