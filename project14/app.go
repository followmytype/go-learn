package main

import (
	"fmt"
)

// 關係運算、邏輯運算、賦值運算
func main() {
	// 關係運算
	// 結果只會是布林值
	// ==
	// !=
	// <
	// >
	// <=
	// >=
	var n1 int = 9
	var n2 int = 7
	fmt.Println("n1:", n1, ", n2:", n2)
	fmt.Println("n1 == n2:", n1 == n2)
	fmt.Println("n1 != n2:", n1 != n2)
	fmt.Println("n1 < n2:", n1 < n2)
	fmt.Println("n1 > n2:", n1 > n2)
	fmt.Println("n1 <= n2:", n1 <= n2)
	fmt.Println("n1 >= n2:", n1 >= n2)

	flag := n1 == n2 // false
	fmt.Println("flag:", flag)

	// 邏輯運算
	// &&, ||, !
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age < 20 || age > 30 {
		fmt.Println("ok2")
	}

	if !(age > 35 && age < 45) {
		fmt.Println("ok3")
	}

	// && 一定要所有條件都為true，所以test()會被執行到
	// 反之如果有一個條件為false，後面的條件都不會再確認
	if age > 20 && test() {
		fmt.Println("ok1....")
	}

	// || 只要一個為true就會自動跳開，所以test不會被執行到
	// 反之只要還沒有true出現，條件就會一直檢查下去
	if age > 20 || test() {
		fmt.Println("ok2....")
	}

	// 賦值運算
	// var i int = 10

	a := 10
	b := 20
	fmt.Printf("a: %v, b: %v\n", a, b)

	a += 5
	b -= 10
	fmt.Printf("a: %v, b: %v\n", a, b)

	// 兩值交換
	a = 3
	b = 2
	fmt.Printf("交換前，a: %v, b: %v\n", a, b)
	a = a + b
	b = a - b // b = (a + b) - b ==> b = a
	a = a - b // a = (a + b) - (a + b) - b ==> a = b
	fmt.Printf("交換後，a: %v, b: %v\n", a, b)
}

func test() bool {
	fmt.Println("test func")

	return true
}
