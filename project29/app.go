package main

import (
	"fmt"
)

// buildin func
func main() {
	number1 := 100
	fmt.Printf("number: %v, number type: %T, number address: %v\n", number1, number1, &number1)

	number2 := new(int)
	// new()是建立一個指標變數
	fmt.Printf("number: %v, number type: %T, number address: %v, number v: %v\n", number2, number2, &number2, *number2)
}
