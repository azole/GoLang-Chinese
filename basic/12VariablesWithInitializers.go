package main

import "fmt"

// var 變數宣告可以給初始值，一個對應一個
var x, y, z int = 1, 2, 3

// 當有給初始值時，就可以省略型別，go 會自動判斷
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)
}
