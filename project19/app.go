package main

import (
	"fmt"
)

// 循環控制
func main() {
	fmt.Println("============")

	// while 與 do...while寫法

	// while寫法
	var i int = 1
	// 先一個無限循環
	for {
		// 設定終止規則
		if i > 10 {
			break
		}
		// 循環要做的行為
		fmt.Println("i =", i)
		// 循環條件的迭代
		i++
	}

	// do...while寫法
	var j int = 1
	// 先一個無限循環
	for {
		// 循環要做的行為，這樣能保證至少會做一次
		fmt.Println("j =", j)
		// 循環條件的迭代
		j++
		// 設定終止規則，將終止條件放在最後
		if j > 10 {
			break
		}
	}

	var num int
	fmt.Println("金字塔層數：")
	fmt.Scanln(&num)
	// *
	// **
	// ***
	// ****
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	//    *
	//   ***
	//  *****
	// *******
	for i := 1; i <= num; i++ {
		starNums := 2*i - 1
		for j := 1; j <= num-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= starNums; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	//    *
	//   * *
	//  *   *
	// *******
	for i := 1; i <= num; i++ {
		starNums := 2*i - 1
		for j := 1; j <= num-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= starNums; k++ {
			if k == 1 || k == starNums || i == num {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
