package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
	// 用來堵塞，不讓主程式結束，不然會看不到結果，
	// 直到 L18 執行了，才會執行這行，然後主程式結束
}
