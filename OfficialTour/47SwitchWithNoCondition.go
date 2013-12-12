// 沒有 condition 的就視同是 switch true
// 這樣的結構很適合用來取代很長的 if-then-else
// 會乾淨很多

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
