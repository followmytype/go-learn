package main

import "fmt"

// golang中，+號演示
func main() {
	// +號用在兩變數中間，如果兩變數為數字，代表相加，如果兩變數為字串，代表字串相連
	var num1 = 1
	var num2 = 2
	var num3 = num1 + num2
	fmt.Println("num3:", num3)

	var str1 = "aaa"
	var str2 = "bbb"
	var result = str1 + str2
	fmt.Println("result:", result)
}
