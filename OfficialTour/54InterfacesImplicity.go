// 一個類型藉由實作那些方法來實作一個介面
// 沒有顯式聲明的必要
//
// 隱式聲明解耦 (decouple) 了實作的 package 跟那些定義了 interface 的 package
// 讓這兩者不互相依賴
// 這樣也鼓勵了明確的介面定義，因為你不用再找到每個實作，然後用新的 interface 名稱標籤它
//
// Package io 定義了 Reader 跟 Writer，你不用這樣做

package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
