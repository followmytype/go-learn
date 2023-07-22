package main

import (
	"fmt"
)

// 小數類型介紹
func main() {
	fmt.Println("============")
	var price1 float32 = 1.23
	fmt.Println("price1:", price1)

	// 小數分為單精度與雙精度
	// float32，單精度，四個字節(byte)，32位元(8*4)
	// float64，雙精度，八個字節(byte)，64位元(8*8)

	// 浮點數在機器中存放形式：符號位+指數位+尾數位（浮點數都是有正負符號的）
	var num1 float32 = -0.02
	var num2 float64 = -7123345.02            // -7.12334502e+06，等於 -1 * 7.12334502 * 10^6
	fmt.Println("num1:", num1, "num2:", num2) // num1: -0.02 num2: -7.12334502e+06

	// 浮點數可能會因為精度造成損失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3:", num3, "num4:", num4) // num3: -123.00009 num4: -123.0000901

	// go的浮點數有固定的範圍和字段長度，不受作業系統影響
	// 預設的float型態為float64
	var num5 = 1.1
	fmt.Printf("num5 type is %T\n", num5) // num5 type is float64

	// 浮點數的表示方法
	// 方法一：十進制
	num6 := 1.23
	num7 := .23                               // 等於 0.23。編寫時建議寫0.23，清楚不賣弄
	fmt.Println("num6:", num6, "num7:", num7) // num6: 1.23 num7: 0.23

	// 方法二：科學表示法
	num8 := 1.23456e4                                          // 等於 1.23456 * 10^4
	num9 := 1.23456E4                                          // 大寫E也沒關係
	num10 := 1.23e-2                                           // 等於 1.23 * 10^-2 等於 1.23 / 10^2
	fmt.Println("num8:", num8, "num9:", num9, "num10:", num10) // num8: 12345.6 num9: 12345.6 num10: 0.0123

	// 通常情況下，推薦使用float64，也就是預設
}
