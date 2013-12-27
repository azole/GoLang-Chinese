// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/03.0.md

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析參數，預設是不會解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("----------------")
	fmt.Fprintf(w, "Hello astaxie!") // 寫入到 w -> 輸出到 client 端
}

func main() {
	http.HandleFunc("/", sayHelloName)       // 設定訪問路由
	err := http.ListenAndServe(":9000", nil) // 設定 port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
