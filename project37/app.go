package main

import (
	"fmt"
)

type Student struct {
	Num int
	Name string
}

func main() {
	// 方法與函數
	// 函式: 單獨一個 test()
	// 方法: 隸屬某個變數 attr.test() 

	std1 := Student{
		Name: "matt",
		Num: 1,
	}
	stud2 := &Student{
		Name: "mina",
		Num: 2,
	}
	fmt.Println(std1, stud2)
}
