// 檢查兩棵 binary tree 的值是否相等
// tree type 已經做好了，請直接使用
// (可以將 tree.go 放在 $GOPATH/tour/tree/tree.go)
//
// 1. 完成 Walk 函式
// 2. 測試 Walk 函式
//
// 函式 tree.New(k) 會建立一個任意結構的二元樹，其值為 k, 2k, 3k, ..., 10k
// 建立一個新的 channel 然後開始遍歷這個樹
// go Walk(tree.New(1), ch)
// 你會看到畫面上印出 1, 2, 3, ..., 10
//
// 3. 利用 Walk 函式實作 Same，Same 是用檢驗兩棵樹的值是否相同
// 4. 測試 Same 函式
// Same(tree.New(1), tree.New(1)) 應該要回傳 true
// Same(tree.New(1), tree.New(2)) 應該要回傳 false

package main

import (
	"fmt"
	"tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//    Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
