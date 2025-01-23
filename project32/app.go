package main

import (
	"fmt"
)

// slice 切片
// 切片是array的引用，在做傳遞時，也是引用傳遞
func main() {
	// 創建切片的第一個方式，引用已經建立好的陣列
	var arr [5]int = [...]int{1, 22, 33, 44, 55}
	// 取index為一到三的元素，但不包含三
	// arr[:3] 前面不寫代表從頭取
	// arr[3:] 後面不寫代表取到尾
	// arr[:] 都不寫代表全部
	slice := arr[1:3]
	fmt.Println("arr = ", arr)
	fmt.Println("slice = ", slice)
	fmt.Println("slice len = ", len(slice))
	fmt.Println("slice cap = ", cap(slice))
	fmt.Printf("arr[1] ptr: %p\n", &arr[1])
	fmt.Printf("slice[0] ptr: %p, slice[0]=%v\n", &slice[0], slice[0])
	// slice本身儲存三個東西
	// 1. 起始地址
	// 2. 長度
	// 3. 容量
	// 所以改變slice內的值或是被引用的arr值都會一起改變 
	fmt.Println("before arr: ", arr, ", slice: ", slice)
	arr[1] = 222
	fmt.Println("after arr: ", arr, ", slice: ", slice)

	// 創建切片的第二個方式，使用make(類型, 長度, 容量)
	// 第三個參數容量選填，需注意如果參數都要寫的話，容量要大於等於長度
	// var makeSlice []int = make([]int, 0, 10)
	makeSlice := make([]int, 5)
	fmt.Println("makeSlice: ", makeSlice, ", len: ", len(makeSlice), ", cap: ", cap(makeSlice))
	var makeSlice1 []int
	fmt.Println("makeSlice1: ", makeSlice1, ", len: ", len(makeSlice1), ", cap: ", cap(makeSlice1))

	makeSlice2 := make([]int, 0)
	// 錯誤，因為切片長度為零
	// makeSlice2[0] = 123
	// 如果要賦值需要這樣寫，
	makeSlice2 = append(makeSlice2, 123) // ok
	// 所以切片在增加值的時候建議都用append

	// 會出現錯誤，因為超過長度了
	// makeSlice[6] = 44

	// 第三種創建方式
	var stringSlice []string = []string{"qq", "ww", "ee"}
	fmt.Println("stringSlice: ", stringSlice, ", len: ", len(stringSlice), ", cap: ", cap(stringSlice))
	fmt.Printf("stringSlice ptr: %p\n", stringSlice)

	// 使用append動態增加元素
	// 可一次增加多個
	stringSlice = append(stringSlice, "rr", "tt")
	fmt.Println("stringSlice: ", stringSlice, ", len: ", len(stringSlice), ", cap: ", cap(stringSlice))
	fmt.Printf("stringSlice ptr: %p\n", stringSlice)
	// 也可增加其他切片
	stringSlice = append(stringSlice, stringSlice...)
	fmt.Println("stringSlice: ", stringSlice, ", len: ", len(stringSlice), ", cap: ", cap(stringSlice))
	fmt.Printf("stringSlice ptr: %p\n", stringSlice)

	// copy, 複製，只限切片使用 
	stringSlice1 := make([]string, 10)
	copy(stringSlice1, stringSlice) // 將後者複製一份到前者去

	// 注意
	var longSlice []int = []int{1,2,3,4,5} // 長度為五
	shortSlice := make([]int, 1) // 長度為一
	copy(shortSlice, longSlice) // 不會錯誤，只會複製一個元素
	fmt.Println("short slice is: ", shortSlice)

	// 切片是引用類型 所以函數內有更動 外面也會一起更動
	fmt.Println("long slice is: ", longSlice)
	changeSlice(longSlice)
	fmt.Println("long slice is: ", longSlice)

	// string也是slice一種
	// 但是string是不可變的
	// 所以不能這樣str[3] = 'C'
	// 如果要更改字串，須將其轉為byte[] or rune[]，rune[]可以處理中文字
	str := "哈囉hello"
	strArr := []rune(str)
	strArr[1] = '欸'
	str = string(strArr)
	fmt.Println(str)
}

func changeSlice(input []int) {
	input[0] = 100
}
