package main

import (
	"fmt"
)

// 字元、字串
// golang中，並沒有專門的字元類型，都是使用byte來做儲存，byte範圍0~255，常見儲存為ASCII碼
// 所有的文字都是利用utf8對應的字元值來表示，所以字元本身是可以做數字運算的，因為他本質就是個數字
func main() {
	fmt.Println("============")

	// 單一字元都是使用單引號表達
	var c1 byte = 'a'
	var c2 byte = '1'
	// 當直接輸出時，就是輸出字元對應的數字碼，參考：https://www.ascii-code.com/
	fmt.Println("c1:", c1, "c2:", c2) // c1: 97 c2: 49
	// 若要轉為文字時，需使用格式化輸出
	fmt.Printf("c1: %c, c2, %c\n", c1, c2)
	// 中文字因為使用unicode，所以要使用int來做輸出（也能使用uint）
	var c3 int = '陳'
	fmt.Printf("c3: %c, c3: %d, c3 type is %T\n", c3, c3, c3) // c3: 陳, c3: 38515, c3 type is int
	var c4 = 'b'
	fmt.Printf("c4 type is %T\n", c4)
	// 可直接對一個變量賦予數字，按照字元的格式化輸出時，會得到對應的字元
	var c5 = 24325
	fmt.Printf("c5 is %c\n", c5) // c5 is 弅

	var n1 = 10
	fmt.Println("n1 =", n1+'a') // n1 = 107

	// 字元儲存到計算機中的步驟
	// 存: 字元->對應值->二進制->存
	// 讀: 二進制->對應值->字元->讀
}
