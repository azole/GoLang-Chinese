// 資料來源：
// 1. https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.6.md
// 2. 官方講解 reflection 的文章 http://blog.golang.org/laws-of-reflection
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	r := reflect.ValueOf(x)
	// 不可修改，會有 panic
	//r.SetFloat(7.1) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println(r)                                               // <float64 Value>
	fmt.Println("type:", r.Type())                               // float64
	fmt.Println("kind is float64:", r.Kind() == reflect.Float64) // true
	fmt.Println("value:", r.Float())                             // 3.4

	// 若想讓 reflect 的值可以修改，要這樣寫：
	v := reflect.ValueOf(&x)
	e := v.Elem()
	e.SetFloat(7.1)
	fmt.Println(x) // 變成 7.1 了
}
