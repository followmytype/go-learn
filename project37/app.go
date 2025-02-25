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

	b := &B{
		A: A{
			Name: "aaaa",
			age: 123,
		},
		Name: "bbbb",
		Gender: "male",
	}
	// 可直接使用繼承的方法，不限於大小寫
	b.A.SayName() // aaaa
	b.A.printAge()
	b.A.age = 222
	b.A.printAge()
	// 也可以忽略中間的繼承結構體寫法
	b.SayName() // bbbb
	b.printAge()
	b.age = 123
	b.printAge()
	// go 會自動在當前結構體內找有沒符合的屬性或方法名稱，如果沒找到才會進去下一個結構體找
	// 就近訪問原則
}

// 繼承
// 在結構體內嵌入另一個結構體，使用匿名方式，即可使用被嵌入的結構體所有屬性與方法，即使是小寫開頭
type A struct {
	Name string
	age int
}

func (a *A) SayName() {
	fmt.Println("hello a name:", a.Name)
}

func (a *A) printAge() {
	fmt.Println("a age is:", a.age)
}

type B struct {
	A
	Gender string
	Name string
}

func (b *B) SayName() {
	fmt.Println("hello b name:", b.Name)
}
