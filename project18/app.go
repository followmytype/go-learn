package main

import (
	"fmt"
)

// 循環控制
func main() {
	fmt.Println("============")
	for i := 1; i <= 10; i++ {
		fmt.Println("hello, i =", i)
	}

	// while寫法
	j := 1
	for j <= 10 {
		fmt.Println("hello, j =", j)
		j++
	}

	// 無限循環
	// for ;; {} // 也可以這樣寫
	k := 1
	for {
		if k <= 10 {
			fmt.Println("無限~")
		} else {
			break
		}
		k++
	}

	var str string = "Hello, World!!哈囉"
	// 如果字串含有中文，會出現錯誤，因為中文字擁有三個字節，所以會出現亂碼
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 將字串轉換為切片即可
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	str = "abc~~okok這裡呦"
	// 此方法可成功顯示中文字
	for index, value := range str {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}
}
