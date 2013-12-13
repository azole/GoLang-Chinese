// 在 select 中，都沒有任何其他的 case 準備好時，就會執行 default case
// Use a default case to try a send or receive without blocking:
// select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 執行結果：
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// tick.
//     .
//     .
// BOOM!
