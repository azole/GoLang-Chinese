// goroutine 是一個由 Go runtime 管理的輕量級執行緒 (lightweight thread)
// go f(x, y, z)
// x, y, z 是由目前的 goroutine 評估，但 f 的執行是在新的 goroutine
//
// Goroutines 是在同一個 address space 中運行，所以存取同一塊共享記憶體 (shared memory) 時必須要同步 (synchronize)
// sync package 提供了一個有用的方法，但你將不會需要用到，因為還有另外的方法(見下一個範例)

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	// 如果是沒有加上 go 只有以下這樣
	// say("world")
	// say("hello")
	// 那執行結果是先印出 5 個 world，再印出 5 個 hello，依序的
}
