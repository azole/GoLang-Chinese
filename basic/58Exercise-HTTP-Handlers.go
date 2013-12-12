// 實現下面的類型，並且為它們定義 ServeHTTP 方法
// 註冊它們以便讓你的 web server 可以處理這些特殊的路徑
//
// 註冊的方式：
// http.Handle("/string", String("I'm a frayed knot."))
// http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
//
// 測試：
// http://localhost:4000/string
// http://localhost:4000/struct

package main

import (
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func main() {
	// your http.Handle calls here
	http.ListenAndServe("localhost:4000", nil)
}
