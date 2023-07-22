package main

import "fmt"

// 定義全域變數
var global1 = 100
var globalName = "AAA"

// 另一方式
var (
	global2     = 200
	globalName2 = "BBB"
)

func main() {
	// 定義變數
	var i int
	// 賦值
	i = 123

	fmt.Println("i =", i)

	// 變數使用方式有三種
	// 第一種，宣告變數，每個變數在宣告後都會有個預設值，若為整數則預設值為0
	var para1 int
	fmt.Println("para1 =", para1)
	// 第二種，根據值直接判定變數型態
	var para2 = "aabb"
	fmt.Println("para2 =", para2)
	// 第三種，省略var，直接賦值，這邊需要用 := 來操作。注意語法直接包含宣告和賦值兩個動作，所以:=左側的變數名稱不能是已經宣告過的
	para3 := 10.12
	fmt.Println("para3 =", para3)

	// 多變量聲明三種方式
	// 第一種，相同類型變數
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)
	// 第二種，宣告時直接賦值
	var n4, name, n5 = 123, "matt", 456
	fmt.Println("n4 =", n4, "name =", name, "n5 =", n5)
	// 第三種，省略var，直接賦值
	n6, animal, n7 := 11, "dog", 22
	fmt.Println("n6 =", n6, "animal =", animal, "n7 =", n7)
}
