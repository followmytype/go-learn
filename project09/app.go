package main

import (
	"fmt"
)

// 基本數據預設默認值
func main() {
	fmt.Println("============")
	var a int     // 0
	var b float32 // 0
	var c float64 // 0
	var d bool    // false
	var e string  // ""
	// a is 0 , b is 0 , c is 0 , d is false , e is
	fmt.Println("a is", a, ", b is", b, ", c is", c, ", d is", d, ", e is", e)
	// a is 0, b is 0.000000, c is 0.000000, d is false, e is
	fmt.Printf("a is %d, b is %f, c is %f, d is %v, e is %v\n", a, b, c, d, e)
	// %v表示按照變數的型態輸出
	// a is 0, b is 0, c is 0, d is false, e is
	fmt.Printf("a is %v, b is %v, c is %v, d is %v, e is %v\n", a, b, c, d, e)
}
