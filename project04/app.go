package main

import (
	"fmt"
	"unsafe"
)

// 整數介紹，int
func main() {
	// golang整數有這幾類
	// int, int8, int16, int32, int64
	// int8代表一個字節byte，8bit，每一bit只有0或1，有符號整數，範圍-128 ~ 127
	// |0 or 1|，代表正負號|0 or 1|0 or 1|0 or 1|0 or 1|0 or 1|0 or 1|0 or 1|
	var num1Int8 = 127
	var num2Int8 = -128
	// var num3Int8 = 128 // 出錯
	// var num4Int8 = -129 // 出錯
	// 同理，int16範圍是 -2^15 ~ 2^15-1
	// 同理，int32範圍是 -2^31 ~ 2^31-1
	// 同理，int64範圍是 -2^63 ~ 2^63-1
	fmt.Println("num1Int8:", num1Int8, "num2Int8:", num2Int8)

	// 無符號整數
	// uint, uint8, uint16, uint32, uint64
	// 最前面一位拿來擺放數值，不做正負號判斷
	// 範圍為 0 ~ 2^(長度 )-1
	var unum1Uint8 uint8 = 0
	var unum2Uint8 uint8 = 255
	fmt.Println("unum1Uint8:", unum1Uint8, "unum2Uint8:", unum2Uint8)

	// 單獨int預設為有符號的，根據作業系統的位元系統而定
	// 32位元系統為四個字節，4 byte，8*4=32，int32，範圍 -2^31 ~ 2^31-1
	// 64位元系統為八個字節，8 byte，8*8=64，int64，範圍 -2^63 ~ 2^63-1

	// uint以此類推
	// 32位元系統為四個字節，4 byte，8*4=32，int32，範圍 0 ~ 2^32-1
	// 64位元系統為八個字節，8 byte，8*8=64，int64，範圍 0 ~ 2^64-1

	// rune
	// 等於int32，不一樣的是他用來表示一個unicode碼

	// byte
	// uint8，0 ~ 255，用來存儲字符

	// 當在宣告整數類型的變數時，默認都是使用int類型的
	// fmt.Printf: 用於格式化輸出
	// %T可以查看變數的類型
	var somenum = 123
	fmt.Printf("somenum type is: %T\n", somenum) // somenum type is: int

	// 如何查看變數使用的字節數
	// 使用到unsafe包的Sizeof()函數
	var onenum int64 = 10
	// onenum value is 10, onenum type is int64, onenum size is 8
	fmt.Printf("onenum value is %d, onenum type is %T, onenum size is %d", onenum, onenum, unsafe.Sizeof(onenum))
}
