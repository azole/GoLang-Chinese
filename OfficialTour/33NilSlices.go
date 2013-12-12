// slice 的零值是 nil
// 一個 nil slice 的長度跟容量都是 0

package main

import "fmt"

func main() {
  var z []int
  fmt.Println(z, len(z), cap(z)) // [] 0 0 
  if z == nil {
    fmt.Println("nil!")
  }
}
