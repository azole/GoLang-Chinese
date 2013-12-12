// 每個 go program 都是由 package 組成的
// 程式會從 package main 開始執行
package main

// 這個範例 import 了 fmt 跟 math/rand 這兩個 packages
// 慣例上，package 的名稱是 import 路徑的最後一個元素
// 例如這邊的 math/rand，去看原始碼，
// 在 /usr/local/go/src/pkg/math/rand 中，會包含數個檔案，
// 但都是以 package rand 作為開頭
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
