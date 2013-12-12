// 在 map 中新增或刪除一個元素: m[key] = elem
// 取回一個元素： elem = m[key]
// 刪除一個元素： delete(m, key)
// 測試一個 key 是否存在，可以用兩個變數的指派
// elem, ok = m[key]
// 如果 map 中有這個 key 存在，ok 會是 true
// 如果不存在，elem 會是該類型的零值，ok 會是 false
// 同樣地，在讀取某個 key 的值時，如果這個 key 不存在，回傳的也會是該類型的零值

package main

import "fmt"

func main() {
  m := make(map[string]int)
  
  m["Answer"] = 42;
  fmt.Println("The value:", m["Answer"]) // 42

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"]) // 48

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"]) // 0

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)  // 0, false
}
