package main

import (
	"fmt"
)

// 函數2
func main() {
	fmt.Println("============")
	// 函式也是一個資料型別
	// 可賦值給一個變數
	a := getSum
	fmt.Printf("a的類型: %T, getSum的類型: %T\n", a, getSum)
	res := a(1, 1)
	fmt.Printf("a(1, 1)回傳%v\n", res)

	// 既然函式可以是一個變數，那他也可以成為一個參數
	myFunc(getSum, 50, 60)
	myFunc(getSub, 50, 60)

	var num1 myInt
	num1 = 10
	var num2 int
	num2 = 10
	fmt.Printf("num1型別為: %T, 值為: %v\nnum2型別為: %T, 值為: %v\n兩者不一樣\n", num1, num1, num2, num2)
	fmt.Printf("需要硬轉換才能讓他們做運算：int(num1) + num2 = %v\n", int(num1) + num2)

	a1 := 10
	a2 := 20
	swap(&a1, &a2)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2

	return sum
}

func getSub(n1 int, n2 int) int {
	sum := n1 - n2

	return sum
}

// 將函式當成參數，第一個參數是函式類型，接收兩個int，回傳int
// 這樣寫的好處是可以不影響原函式的內容下做其他的包裝
// 或是定義好介面讓開發者遵循
func myFunc(somefunc func(n1, n2 int) int, n1, n2 int) int {
	fmt.Printf("n1: %v, n2: %v\n", n1, n2)
	res := somefunc(n1, n2)
	fmt.Printf("處理結果: %v\n", res)
	return res
}

// golang可以自定義資料型別
// 這個寫法是定義一個myInt資料型別
// 但是他的本質上是int
type myInt int

// 同理也能將函式定義為新的型別
type twoNumFunc func(n1, n2 int) int

func anyTwoNum(fn twoNumFunc, n1, n2 int) int {
	fmt.Printf("第一個參數的型別為: %T\n", fn)
	res := fn(n1, n2)
	fmt.Printf("處理結果為: %v\n", res)
	return res
}

func swap(n1, n2 *int) {
	t := n1
	n1 = n2
	n2 = t
}
