// http 通過任何實作了 http.Handler 的值來處理 HTTP requests
//
// package http
//
// type Handler interface {
//   ServeHTTP(w ResponseWriter, r *Request)
// }
// 下面的範例中，類型 Hello 就實作了 http.Handler
// 執行後，拜訪 http://localhost:4000

package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
