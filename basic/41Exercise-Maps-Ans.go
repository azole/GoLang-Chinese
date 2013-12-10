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
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

// 最後輸出的結果：
// PASS
//  f("I am learning Go!") =
//   map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") =
//   map[string]int{"brown":1, "over":1, "lazy":1, "dog.":1, "The":1, "quick":1, "fox":1, "jumped":1, "the":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") =
//   map[string]int{"I":2, "ate":2, "a":1, "donut.":2, "Then":1, "another":1}
// PASS
//  f("A man a plan a canal panama.") =
//   map[string]int{"A":1, "man":1, "a":2, "plan":1, "canal":1, "panama.":1}
