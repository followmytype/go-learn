package main

import (
	"fmt"
)

// 指標複習、使用者輸入
func main() {
	fmt.Println("============")
	var i int = 100
	var iPtr *int = &i
	fmt.Printf("i 的位址是%v\n", iPtr)
	fmt.Printf("i 的值是%v\n", *iPtr)

	var name string
	var age byte
	var sal int
	var pass bool

	fmt.Println("姓名：")
	fmt.Scanln(&name) // 這邊記得放變數的地址
	fmt.Println("年齡：")
	fmt.Scanln(&age)
	fmt.Println("薪水：")
	fmt.Scanln(&sal)

	fmt.Printf("姓名：%v, 年齡：%v, 薪水：%v\n", name, age, sal)

	fmt.Println("輸入姓名、年齡、薪水、是否畢業，用空格隔開")
	fmt.Scanf("%s %d %d %t", &name, &age, &sal, &pass)
	fmt.Printf("姓名：%v, 年齡：%v, 薪水：%v, 是否畢業：%v\n", name, age, sal, pass)
}
