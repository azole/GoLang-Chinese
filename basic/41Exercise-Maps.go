// 實作 WordCount
// 這個函式會回傳一個 map，這個 map 中會存有 s 中每個字出現的次數
// wc.Test 會執行一個測試函式，並且印出成功或失敗
// 例如：
// FAIL
//  f("I am learning Go!") =
//   map[string]int{"x":1}
//  want:
//   map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
//
// 提示：可以使用 strings.Fields 這個函式

package main

import (
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
