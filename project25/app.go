package main

import (
	"fmt"
)

// defer
// 代表延遲執行的意思，當他所在的作用域結束時才會執行
// 執行順序為先進後出的規則
func main() {
	fmt.Println("hello, im 1")
	defer fmt.Println("hello, im 2")
	defer fmt.Println("hello, im 3")
	fmt.Println("hello, im 4")
	n1 := 3
	// 當變數值放入defer時，內容已經被鎖定了，後續改變內容值時不影響defer輸出
	defer fmt.Println("defer do, n1 is", n1) // defer do, n1 is 3
	n1 = 4
	fmt.Println("n1 is", n1)

	fmt.Println("someFunc():", someFunc())
}

func someFunc() int {
	n1 := 10
	n2 := 20
	defer fmt.Println("n1:", n1)
	defer fmt.Println("n2:", n2)

	fmt.Println("n1 + n2 =", n1+n2)
	return n1 + n2
}
