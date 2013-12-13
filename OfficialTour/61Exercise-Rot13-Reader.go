// 一個很常見的模式是：一個 io.Reader 包著另外一個 io.Reader，用一些方法來修改流 (stream)
//
// 例如，gzip.NewReader 裡頭就有包著 io.Reader，但實作了一個解壓縮的資料流
//
// 這個練習是請你實作 rot13Reader，修改資料流使其符合 ROT13 (http://en.wikipedia.org/wiki/ROT13)
// 目前已經幫你把 rot13Reader 的類型做好了，請實作 Read 這個方法 (io.Reader的介面)

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
