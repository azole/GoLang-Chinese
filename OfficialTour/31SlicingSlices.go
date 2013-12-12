// slices 可以被重新分割 (slice)，建立一個新的 slice，但仍指向同一個陣列
// 見下方譯者自己新增的測試，s 是 p slice 出來的，
// 修改了 s 內的值後，p 也會被修改
// 分割 (slice) 的語法為 s[lo:hi]
// 會把從 lo 到 hi-1 的元素分割出來
// 所以 s[lo:lo] 這樣會是空的
// 而 s[lo:lo+1] 會有一個元素

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)           // p ==  [2 3 5 7 11 13]
	fmt.Println("p[1:4] == ", p[1:4]) // p[1:4] ==  [3 5 7]

	// 如果 low 沒有寫，就是從 0 開始
	fmt.Println("p[:3] == ", p[:3]) // p[:3] ==  [2 3 5]

	// 如果high 沒有寫，就是直到原本 slice 的長度 len(p)
	fmt.Println("p[4:] == ", p[4:]) // p[4:] ==  [11 13]

	// 譯者自己新增的測試，測試指向同一個陣列
	s := p[1:4]
	s[0] = 1000
	fmt.Println("s == ", s) // s ==  [1000 5 7]
	// 原來的 p[1] 也被修改成 1000 了
	fmt.Println("p == ", p) // p ==  [2 1000 5 7 11 13]

}
