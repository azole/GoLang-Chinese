/* Generator: 回傳 channel 的函式
 *
 * Channels as a handle on a service
 * 函式回傳一個 channel，讓我可以跟它提供的服務溝通
 */

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

func main() {
	// c := boring("boring!")
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %q\n", <-c)
	// }

	azole := boring("azole")
	kate := boring("kate")
	for i := 0; i < 5; i++ {
		fmt.Println(<-azole)
		fmt.Println(<-kate)
	}
	// 這樣做一定是 azole, kate, azole, kate,... 的順序

	fmt.Println("You're boring: I'm leaving.")
}
