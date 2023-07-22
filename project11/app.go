package main

import (
	"fmt"
	// "strconv"
)

// 指標
func main() {
	fmt.Println("============")
	var i int = 10
	fmt.Printf("i 值為 %v\n", i)
	fmt.Printf("i 的址為 %v\n", &i)

	var ptr *int = &i
	// ptr 變數名
	// ptr 類型為 *int，int的指標
	// ptr 值為 &i，i的地址
	fmt.Printf("ptr 值為 %v\n", ptr)

	fmt.Printf("ptr 自己的址: %v\n", &ptr)
	fmt.Printf("ptr 的址所儲存的值: %v\n", *ptr)

	*ptr = 20
	fmt.Printf("i 值為 %v\n", i)

	// 值類型，都有對應的指標類型，int -> *int, float64 -> *float64
	// 值類型包括 int, float, bool, string, 數組, 結構struct

	// 引用類型，指標，slice切片，map，管道channel，interface
}
