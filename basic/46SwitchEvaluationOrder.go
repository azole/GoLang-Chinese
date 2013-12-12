// switch cases 評估的順序是由上到下，直到有一個 case 符合。
//
// 例如
//
// switch i {
//   case 0:
//   case f():
// }
// 如果 i==0 的話，f 就不會被呼叫到

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
