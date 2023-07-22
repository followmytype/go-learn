package main

import (
	"fmt"
)

// 布林值
// 布林值佔一個字節（1 byte, 8 bits
func main() {
	fmt.Println("============")
	var a = true
	fmt.Println("a is", a)

	var b = false
	fmt.Println("b is", b)
	// 布林值只能是true or false，不能是0 or 1等其他數值
	// b = 0 // 出錯
}
