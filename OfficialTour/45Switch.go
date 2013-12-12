// 你大概可以猜到 switch 的語法是怎麼樣子了
// case 能自動 break，除非它以 fallthrough 結尾
// --> 這一點我覺得很有趣，
//     switch fall through 的功能，在 C++ 之類的語言是自動的，
//     有時候會有很方便的用途，但是當然也會造成一些問題，
//     例如沒有要 fall through，但工程師忘了寫 break；
//     或是這樣會破壞程式的流程，不易於維護等等原因。
//     所以即使有自動 fall through 的程式語言，大家都會建議一定要加 break，
//     或是最好不要用到自動 fall through，像是 Crockford 的程式規範 (js) 就禁止使用。
//     但我一直覺得這樣有點矯枉過正，因為有些時候用 fall through 真的很好用。
//     Nicbolas C. Zakas 在 Maintainable Javascript 一書中有提到，
//     可以用，但建議在用自動 fall through 功能時，要利用註解明確的說明。
//     這個說法我就很能接受，但畢竟要利用註解說明，這得靠工程師的自律，
//     現在看到 go 的做法就超欣賞的，關閉這個功能以免犯錯，
//     要用的時候加關鍵字，幾乎等同于強迫工程師要註解說明一樣，
//     這樣真的是兼顧了正確性與可維護性。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough   // 加上這個關鍵字，當這個 case 執行完時，下一個 case 也會被執行
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.", os)
	}
}
