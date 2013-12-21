// slice 會指向一個 array，並且包含一個長度
// []T 是指 slice 中的元素是 type T

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

/*
p ==  [2 3 5 7 11 13]
p[0] == 2
p[1] == 3
p[2] == 5
p[3] == 7
p[4] == 11
p[5] == 13
*/

// 補充說明：
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.2.md
//
// go 另外還有一種資料結構：array
// var arr [n]type
// 長度為 n，型別為 type 的陣列，[3]int 跟 [4]int 是不同的陣列類型，
// 當我們把一個陣列傳入函數時，傳遞的是複製的值，而不是 pointer，
// slice 傳遞的才是 pointer
//
// 其他宣告方式：
// a := [3]int{1, 2, 3}
// b := [10]int{1, 2, 3}  // 長度為 10，但只有前面三個有初始化
// c := [...]int{4, 5, 6} // 長度省略，go 會根據元素個數來計算長度
//
// 2 維陣列：
// doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
// 或是
// easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
//
// slice 可以說是一種動態的 array，它其實是一種引用類型，底層才指向一個 array。
// slice 的宣告跟 array 很像，但是不用指明長度
//
