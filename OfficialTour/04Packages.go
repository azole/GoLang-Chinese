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

// 補充說明：
// 來源：https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/02.1.md
//
// package <pkgName>
// 如果 pkgName 為 main，表示此為一個可以獨立運作的 package，編譯後會產生可執行的檔案。
// 除了 main 以外，其他名稱的 package 都是產生一個 .a 的檔案，會放在 $GOPATH/pkg 中。
// --> 每一個可以獨立運作的 go 程式，必定包含一個 package main，
//     而這個 main 中，也必定包含一個 main 函式，這個 main 函式沒有參數也沒有回傳值。
//
// 這樣就達成了模組化，也有了模組化的可重用性。
