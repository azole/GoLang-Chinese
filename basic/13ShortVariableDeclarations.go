package main

import "fmt"

func main() {
	// 在函式內，變數的宣告方式與在函式外差不多
	var x, y, z int = 1, 2, 3
	// 不同的地方在於，函式內的隱含型別宣告可以用 := 來取代
	//var c, python, java = true, false, "no!"
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}
