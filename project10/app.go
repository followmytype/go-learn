package main

import (
	"fmt"
	"strconv"
)

// 變量類型轉換
// golang中的變數類型轉換全部是硬轉換
func main() {
	fmt.Println("============")
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	// i is 100, n1 is 100, n2 is 100, n3 is 100
	fmt.Printf("i is %v, n1 is %v, n2 is %v, n3 is %v\n", i, n1, n2, n3)

	// 被轉換的數值依然是原先的變數型態，以上面的例子來說，i的資料型態依舊是int32
	// i type is int32
	fmt.Printf("i type is %T\n", i)

	// overflow 溢出，大轉小不會出錯，只是會溢出
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	// num2 is 63
	fmt.Printf("num2 is %d\n", num2)
	// 原因：999999的二進值為1111 0100 0010 0011 1111
	// 轉換成int8只取後面八碼，所以是0011 1111，轉換成十進制即為63

	var someVar1 int32 = 100
	var someVar2 int64
	var someVar3 int8

	// someVar2 = someVar1 + 20 // 會出錯，因為等號右邊是int32的運算結果，當要給int64時會因為沒有轉換而報錯
	// someVar3 = someVar1 + 20
	someVar2 = int64(someVar1) + 20 // 成功
	someVar3 = int8(someVar1) + 20  // 成功
	fmt.Println("someVar2 is", someVar2, ", someVar3 is", someVar3)

	someVar1 = 1
	var var4 int8
	// var var5 int8
	var4 = int8(someVar1) + 127 // 編譯會成功，但因為int8的正數最大值為127，所以會溢出，變成-128
	// var5 = int8(someVar1) + 128 // 編譯會失敗，因為var5是int8，最大值只能到127，當右邊有超過127的數值時，會直接不給編譯
	fmt.Println("var4 is", var4)

	// 轉字串(string)
	// 1. fmt.Sprintf(): string
	// https://studygolang.com/pkgdoc 裡的fmt包有寫百分比符號的各種解釋
	var someNum int = 99
	var someFloat float64 = 0.1234
	var someBool bool = false
	var someChar byte = 'G'
	var str string

	str = fmt.Sprintf("%d", someNum)
	fmt.Printf("str type is %T, str is %q\n", str, str)

	str = fmt.Sprintf("%f", someFloat)
	fmt.Printf("str type is %T, str is %q\n", str, str)

	str = fmt.Sprintf("%t", someBool)
	fmt.Printf("str type is %T, str is %q\n", str, str)

	str = fmt.Sprintf("%c", someChar)
	fmt.Printf("str type is %T, str is %q\n", str, str)

	var someNum1 int = 999
	var someFloat1 float64 = 123.345567
	var someBool1 bool = true

	// FormatInt(整數 需int64, 輸出進制)
	str = strconv.FormatInt(int64(someNum1), 10)
	fmt.Printf("str type is %T, str is %q\n", str, str) // str type is string, str is "999"
	// Itoa(整數 int)
	// str = strconv.Itoa(someNum1)

	// FormatFloat(小數 需float64, 格式, 精度, 表示f的來源類型)
	str = strconv.FormatFloat(someFloat1, 'f', 10, 64)
	fmt.Printf("str type is %T, str is %q\n", str, str) // str type is string, str is "123.3455670000"
	str = strconv.FormatFloat(someFloat1, 'f', 10, 32)
	fmt.Printf("str type is %T, str is %q\n", str, str) // str type is string, str is "123.3455657959"

	// FormatBool(布林值)
	str = strconv.FormatBool(someBool1)
	fmt.Printf("str type is %T, str is %q\n", str, str) // str type is string, str is "true"

	// 字串轉其他資料型態
	var str1 string = "true"
	var b bool
	// ParseBool會回傳兩個值(value bool, err error)，利用底線_代表忽略錯誤值
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type is %T, b is %v\n", b, b)

	var str2 string = "130"
	var someInt int64
	// ParseInt(轉換值, 進制, 返回值類型 0->int 32->int32 64->int64)
	// 第二個參數為進制判斷，當為0時，go會自動判斷字串為哪個進制，"0b"開頭為二進制，"0"or"0o"開頭為八進制，"0x"為十六，其餘皆為十進制
	// 第三參數定義回傳指定的位數結果，但實際上回傳值的類型為int64，所以必須先找個64的變數接值，然後再轉為其他位數
	// 以下範例是將128轉為8位元的整數，因為8位元整數最高只到127，所以在轉換時會回傳127
	// 但他還是將127以int64的型態回傳
	someInt, _ = strconv.ParseInt(str2, 10, 8)
	fmt.Printf("someInt type is %T, someInt is %v\n", someInt, someInt) //someInt type is int64, someInt is 127

	// ParseFloat 使用同上
	// ParseFloat(轉換值, 返回值類型) (f float64, err error)
	var str3 string = "123.456"
	var someFloat2 float64
	someFloat2, _ = strconv.ParseFloat(str3, 64)
	// someFloat2 type is float64, someFloat2 is 123.45600128173828 -> 32的結果
	// someFloat2 type is float64, someFloat2 is 123.456 -> 64的結果
	fmt.Printf("someFloat2 type is %T, someFloat2 is %v\n", someFloat2, someFloat2)

	// 字串轉為其他的資料型態時，需注意字串是否為有效數值，例如hello就不能轉為整數
	// 若為無效字串，則會轉為目標型態的預設值，hello -> int = 0
}
