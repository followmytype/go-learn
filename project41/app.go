package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point = Point{1,2}
	a = point // ok
	var b Point
	// b = a // not ok
	b = a.(Point) // 將一個空接口轉移到另一個對應的類型要這樣寫
	fmt.Println(b)

	var x interface{}
	var floatNum float32 = 1.1
	x = floatNum
	y := x.(float32) // 如果轉換成float64則會爆panic
	fmt.Printf("y type is %T, y value is %v\n", y, y)
	// 類型轉換時可以帶個檢查寫法
	// y, ok := x.(float32)
	// if ok {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// 簡潔寫法
	// if y, ok := x.(float32); ok {
	// 	fmt.Println("ok")
	// }

	// 實際應用
	var MyInterfaceArr [3]MyInterface
	MyInterfaceArr[0] = StructOne{}
	MyInterfaceArr[1] = StructOne{}
	MyInterfaceArr[2] = StructTwo{}
	for _, v := range MyInterfaceArr {
		WorkingMyInterface(v)
	}
}

type MyInterface interface {
	Func1()
	Func2()
}

type StructOne struct {}
func (s StructOne) Func1() {
	fmt.Println("StructOne Func1")
}
func (s StructOne) Func2() {
	fmt.Println("StructOne Func2")
}

type StructTwo struct {}
func (s StructTwo) Func1() {
	fmt.Println("StructTwo Func1()")
}
func (s StructTwo) Func2() {
	fmt.Println("StructTwo Func2()")
}
func (s StructTwo) Func3() {
	fmt.Println("StructTwo Func3()")
}

func WorkingMyInterface(mi MyInterface) {
	mi.Func1()
	mi.Func2()
	if st, ok := mi.(StructTwo); ok {
		st.Func3()
	}
}
