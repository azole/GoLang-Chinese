// Slices 可以透過 make 建立，make 會分配一個初始值為 0 的陣列，並回傳一個指向該陣列的 slice。
// 範例： a := make([]int 5) // len(a) = 5
// 如果要特別聲明 capacity (容量)，可以用第三個參數特別聲明
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5 --> 長度為0 但容量為5
// b = b[:cap(b)]         // len(b)=5, cap(b)=5 --> 所以可以擴展其長度        
// b = b[1:]              // len(b)=4, cap(b)=4
// len 是長度，capacity 是容量
 
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5) // 長度是 0, capacity 是 5
	printSlice("b", b)     // b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0]

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
