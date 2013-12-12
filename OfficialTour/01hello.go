// 可以利用這個小程式來測試環境安裝是否已經完成、指令是否能正確執行
// 直接編譯與執行  go run 01hello.go
// 編譯 go build 01hello.go
// 編譯後會產生一個 01hello.exe (in Windows)，這個 exe 檔可以直接被執行

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
