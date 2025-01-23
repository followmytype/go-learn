package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// array
func main() {
	// 定義時已經有預設值了
	var myArr [3]int
	fmt.Println("myArr 值", myArr)
	fmt.Printf("myArr 址：%p\n", &myArr)
	for i, v := range myArr {
		fmt.Printf("%v\n", v) // 0
		fmt.Printf("%p - %p\n", &v, &myArr[i])
	}

	var intArr [3]int
	for i := range intArr {
		fmt.Printf("輸入第%d個值: ", i+1)
		fmt.Scanln(&intArr[i])
	}
	fmt.Println(intArr)

	var myArr1 [3]int = [3]int{1, 23, 34}
	var myArr2 = [3]int{1, 435, 123}
	var myArr3 = [...]int{11, 22, 33}
	fmt.Println(myArr1, myArr2, myArr3)

	var names = [...]string{1: "tom", 0: "mary", 3: "jerry"}
	fmt.Println(names)

	fmt.Printf("before test01, myArr1[0]:%d\n", myArr1[0])
	test01(myArr1)
	fmt.Printf("after test01, myArr1[0]:%d\n", myArr1[0])

	fmt.Printf("before test02, myArr1[0]:%d\n", myArr1[0])
	test02(&myArr1)
	fmt.Printf("after test02, myArr1[0]:%d\n", myArr1[0])


	printChar()
	fmt.Println()

	sum := 0
	for _, v := range myArr1 {
		sum += v
	}
	fmt.Printf("myArr1總和: %d, 平均值: %v\n", sum, float64(sum) / float64(len(myArr1)))

	var myArr4 [5]int64
	for i := 0; i < len(myArr4); i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Println("n=", n)
		myArr4[i] = n.Int64()
	}
	fmt.Println(myArr4)
}

// 在go中，array的長度視同型態的一部份，所以[3]int與[4]int是兩種不同的類型
func test01(arr [3]int)  {
	// 陣列在function裡傳遞時，為值傳遞，意即更改array內容，不會影響到外部的變數內容
	arr[0] = 123
}

// 指標應用
// 使用指標效率會比較高，因為所用的空間只有指標變數的空間，當使用值傳遞時，就會需要複製傳遞值
func test02(arr *[3]int)  {
	// arr 目前是指標變數，所以要先取值，(*arr)這樣就代表array了
	(*arr)[0] = 123
}

func printChar()  {
	var char [26]byte
	for i := 0; i < len(char); i++ {
		char[i] = 'A' + byte(i)
	}

	for _, v := range char {
		fmt.Printf("%c ", v)
	}
}
