package main

import (
	"fmt"
	"go-learn/project13/tool"
)

// 算術運算
func main() {
	fmt.Println(tool.Add(1, 2))
	fmt.Println(tool.Subtract(5, 2))

	// 除法重點
	// 除法中參與的數值類型為整數的話，出來的結果只會有整數部分，不會有小數，不會四捨五入
	fmt.Println(10 / 4) // 2

	var n1 float32 = 10 / 4 // 2
	fmt.Println(n1)         // 2

	// 若需要算出２.5，則參與計算的成員必須是浮點數
	var n2 float32 = 10.0 / 4
	fmt.Println(n2) // 2.5
	var n3 float32 = 10 / 4.0
	fmt.Println(n3) // 2.5

	// 取餘重點
	// 公式: a % b = a - a / b * b
	fmt.Println("10 % 3 =", 10%3)     // 1
	fmt.Println("-10 % 3 =", -10%3)   // -1
	fmt.Println("10 % -3 =", 10%-3)   // 1
	fmt.Println("-10 % -3 =", -10%-3) // -1

	// ++ --
	var i int = 10
	fmt.Println("i =", i) // 10
	i++
	fmt.Println("i++ =", i) // 11
	i--
	fmt.Println("i-- =", i) // 10
	// 自增自減只能當作一個獨立語句使用
	// b := a++ // 錯誤
	// b := a-- // 錯誤
	// go沒有++a, --a這種

	// 練習
	// 97天等於幾個星期幾天
	var days int = 97
	var week int = days / 7
	var leftday int = days % 7
	fmt.Printf("97天等於%d個星期又%d天\n", week, leftday)

	// 華氏溫度轉攝氏溫度
	// 5/9*(華氏溫度-100)
	var f float32 = 134.2               // 華氏溫度
	var c float32 = 5.0 / 9 * (f - 100) // 攝氏溫度，注意這邊的5一定要寫5.0，不然5/9會變成0，計算就會出問題
	fmt.Printf("華氏溫度%v等於攝氏溫度%v\n", f, c)
}
