package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析參數，對於 POST 則是解析 request body
	// 注意：預設不會去解析，所以一定要呼叫 ParseForm，下面的程式才能執行
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("-----------------------")
	fmt.Fprintf(w, "Hello astaxie!") // 寫入到 w -> 輸出到 client 端
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 獲取請求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // 記得要自己呼叫解析
		// 是 login 的 request，執行 login 的動作
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		// 也可以用 r.FormValue("username") 來取得資料
		// 用這種方法的話，可以不用先叫用 r.ParseForm
		// 它會自動呼叫
		// r.FormValue 只會回傳同名參數中的第一個
		// 若參數不存在，則返回空字串
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 設定路由
	http.HandleFunc("/login", login)         // 設定登入的路由
	err := http.ListenAndServe(":9000", nil) // 設定 port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
