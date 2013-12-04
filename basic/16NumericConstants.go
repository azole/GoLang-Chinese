// TODO: 待研究

package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//fmt.Println(Big) // 編譯時會出現 constant 1267650600228229401496703205376 overflows int
	fmt.Println(Small)          // 2
	fmt.Println(needInt(Small)) // 21
	//fmt.Println(needInt(Big))     // 編譯時會出現 constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
}
