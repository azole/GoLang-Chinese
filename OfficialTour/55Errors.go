// An error is anything that can describe itself as an error string.
// The idea is captured by the predefined, built-in interface type, error,
// with its single method, Error, returning a string:
//
// type error interface {
//   Error() string
// }
//
// fmt 印東西的函式自動會知道要去呼叫這個 Error() 來印出一個錯誤
//
// c9s補充：
// func (self *error) Writer(w io.Writer) {
//	 fmt.Fprint(w, *self)
// }
// func (self *error) String(){
//   return *self
// }

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// go 沒有一般的異常機制，但有 panic 跟 recover 機制，但這應該要很少很少被用到。
// 可參考：
// 1. http://blog.golang.org/defer-panic-and-recover
// 2. https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.3.md
