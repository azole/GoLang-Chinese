// 1. 常數地宣告方式跟變數一樣，但關鍵字是 const
// 2. 常數可以是自源、字串、布林值或數值
// 3. 常數不可以用 := 宣告

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "world"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
