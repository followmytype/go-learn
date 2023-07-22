package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 循環控制
func main() {
	fmt.Println("============")

	rand.Seed(time.Now().UnixNano())
	// 生成 [0,n) 的數字，不包括n
	// 此寫法會生成 1 ~ 100 數字
	n := rand.Intn(100) + 1
	fmt.Println(n)

	times := 0
	for {
		n := rand.Intn(100) + 1
		times++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99一共用了%d次\n", times)


	// for 迴圈可以賦予標籤，作為指定break, continue的用途
	label1:
	for i := 0; i <= 10; i++ {
		fmt.Println("i :", i)
		// label2:
		for j := 0; j <= 5; j ++ {
			fmt.Printf("j : %d \t", j)
			if 7 == i {
				break label1
			}
		}
		fmt.Println()
	}
}
