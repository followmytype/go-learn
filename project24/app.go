package main

import (
	"fmt"
	"strings"
)

// 全域匿名函數
var Superfunc = func(n1 int, n2 int) int {
	return n1 * n2
}

// 閉包，可使用外部的n值，與他形成一個整體
// 這個整體可視為一個類別，每當執行一次，都會保留修改後的ｎ值
func CloseFunc() func(int) int {
	n := 10
	fmt.Println("現在閉包裡的n值為：", n)
	return func(x int) int {
		n += x
		return n
	}
}

// 匿名函數
func main() {
	// 基本定義，在宣告時不用寫定義函數名稱，在函數末端加上參數代表直接執行
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("result:", result)

	// 匿名函數可作為變數宣告
	// a不能理解為函數名，a還是一個變數
	somefunc := func(n1 int, n2 int) int {
		return n1 - n2
	}

	result = somefunc(30, 15)

	fmt.Println("result:", result)

	result = Superfunc(2, 5)

	fmt.Println("result:", result)

	close := CloseFunc()
	fmt.Println("第一次執行close", close(10)) // 20
	fmt.Println("第二次執行close", close(10)) // 30
	fmt.Println("第三次執行close", close(10)) // 40

	jpgSuffix := makeSuffix(".jpg")
	fmt.Println("'a.jpg' =>", jpgSuffix("a.jpg"))
	fmt.Println("'a' =>", jpgSuffix("a"))
	pngSuffix := makeSuffix(".png")
	fmt.Println("'a.png' =>", pngSuffix("a.png"))
	fmt.Println("'a' =>", pngSuffix("a"))
}

// 閉包應用：
// 製作一個必包涵式，可傳入自定義的後綴參數後，回傳一個判斷後綴的函式
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}
// 閉包的好處為，傳統函式在調用時，內部的變數都會初始化
// 並不會保留改動後的值，而閉包能保留並持續應用，可視情況需求選擇閉包或是傳統函式
