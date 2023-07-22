package main

import (
	"fmt"
)

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

// 函數
func main() {
	fmt.Println("============")

	result1, result2 := getSumAndSub(1, 3)
	fmt.Printf("result1: %v, result2: %v\n", result1, result2)
}
