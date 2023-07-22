package main

import (
	"fmt"
)

// 字串
func main() {
	fmt.Println("============")
	var address string = "新北市中和區 235 hello world"
	fmt.Println(address)

	// 字串一旦賦值了，就不可變動
	// var str = "hello"
	// str[0] = 'g' // 無法修改內容，會出錯

	// 字串的表示符號有兩種，一種是雙引號，一種為反引號
	// 雙引號能轉換轉義字元
	str1 := "gg \nininder"
	// gg
	// ininder
	fmt.Println(str1)

	// 反引號能輸出字串的真實內容
	str2 := `gg \n ininder`
	// gg \n ininder
	fmt.Println(str2)

	// 字串拼接
	str3 := "hello " + "world "
	str3 += "haha"
	fmt.Println(str3)

	// 當多個字串要拼接時，可以換行，但要注意+號需放在上一行的尾端
	// 理由：因為golang沒有利用分號去做一段敘述的結束標示，所以如果+號不放在尾端，golang就會以為這行的敘述結束了，但因為下一行的開頭是+號，所以會出錯(我猜的)
	str4 := "aa" +
		"bb" +
		"cc" +
		"dd"
	fmt.Println(str4)
}
