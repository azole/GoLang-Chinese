package main

import "fmt"

// var 宣告了一串變數，注意，型別也是放在後面
var x, y, z int
var c, python, java bool

func main() {
	fmt.Println(x, y, z, c, python, java) // 0 0 0 false false
}

// 補充說明：
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.2.md
//
// 這邊把各種方式都放在一起，後面都會看到範例程式碼。
//
// 1. 使用 var 定義變數：定義類型都是 type 的變數
//    var variableName type
//
// 2. 一次定義多個變數：定義三個類型都是 type 的變數
//    var vname1, vname2, vname3 type
//
// 3. 定義變量並且初始化：定義型都是 type 的變數，其值為 value
//    var variableName type = value
//
// 4. 同時初始化多個變數：
//    var vname1, vname2, vname3 type = v1, v2, v3
//
// 5. 定義與初始化多個變數，但不需要型別，go 會根據初始化的值來定義其型別
//    var vname1, vname2, vname3 = v1, v2, v3
//
// 6. 省略 var，但這只能在函數內部使用，全局變數還是必須用 var 來定義
//    vname1, vname2, vname3 := v1, v2, v3
//
