package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // 回傳一個只允許接收的channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("azole"), boring("kate"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	// 這邊就不一定是 azole, kate, azole, kate,... 這樣依序輪流的順序了
	// 有可能會連續都出現其中一個之類的
	// 可以看 01Multiplexing.png
	fmt.Println("You're boring: I'm leaving.")
}
