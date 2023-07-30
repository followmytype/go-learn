package main

import (
	"fmt"
	"reflect"
)

// 函數
// 作用域：同一個包底下的任何地方都可調用，也就是同一個目錄底下的文件內都可調用
// 如果函式名稱為小寫，則只能在該包底下使用，大寫則可在其他包的區域內調用
func main() {
	fmt.Println("============")

	result1, result2 := getSumAndSub(1, 3)
	fmt.Printf("result1: %v, result2: %v\n", result1, result2)
	println("============")
	num1 := 123
	println("num1 is", num1)
	addNumOne(&num1)
	println("num1 is", num1)
	println("============")
	num2 := 456
	println("num2 is", num2)
	num2 = addNumOneBetter(num2)
	println("num2 is", num2)
	println("============")
	// 長參數呼叫方式
	longParam(11, 22, 33, 44, 55)
	intArr := []int{1, 2, 3, 4, 5}
	// 也可以用陣列當作參數，單需要展開
	longParam(intArr...)
}

// go 的參數如果都是一樣的型別，則可以只保留最後一個參數
// func getSumAndSub(n1, n2 int) (int, int) {
//     ...
// }

// go 返回值型別寫在函式最後，可以返回多個值，若只有一個的話可不用小括弧
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

// 返回值可以命名，這樣在函式底部只需要return就好
// 他相當於先宣告一個回傳的變數
// 此方法建議只用於短小函式，大型函示會影響閱讀
func returnNamedVar() (num int) {
    num = 123
    return
}

// 指標參數，可直接利用函式去更改變數的值
func addNumOne(n1 *int) {
	*n1 += 1
}

// 根據clean code的看法，上面這個寫法稱為輸出型參數
// 參數在函式內操作後會影響到外部的變數值
// 使得參數難以閱讀和維護，推薦使用回傳值來做變數的調整
// 上述程式碼會更改為如下
func addNumOneBetter(n1 int) int {
	n1++
	return n1
}

// 長參數：如果函式的參數會有不定數量的輸入，可使用以下寫法
func longParam(nums ...int) {
	for _, v := range nums {
		println(v)
	}
}

// 如果參數型別不定的話可使用此寫法
func longAnyParam(params ...interface{}) {
	for _, param := range params {
		switch reflect.TypeOf(param).Kind() {
		case reflect.Int:
			fmt.Println(param, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", param)
		case reflect.Array:
			fmt.Println(param, "is an array type.")
		default:
			fmt.Println(param, "is an unknown type.")
		}
	}
}
