package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("catch err:", err)
		}
		// if err := recover(); err != nil { // 相同寫法
		// }
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
	fmt.Println("hello")
}

func test2()  {
	err := errorReturn()
	panic(err)
}

func errorReturn() error {
	return errors.New("hello error")
}

// 錯誤處理
func main() {
	test()

	fmt.Println("still print")

	test2()
	fmt.Println("not print")
}
