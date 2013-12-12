// An error is anything that can describe itself as an error string.
// The idea is captured by the predefined, built-in interface type, error,
// with its single method, Error, returning a string:
//
// type error interface {
//   Error() string
// }
//
// fmt 印東西的函式自動會知道要去呼叫這個 Error() 來印出一個錯誤

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
