package main

import (
	"fmt"
)

// 流程控制
func main() {
	fmt.Println("============")
	var age byte
	fmt.Println("請輸入年齡：")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你未成年")
	}

	if price := 100; price >= 150 { // if 內可聲明新的變數，但作用域只在這個區塊內
		fmt.Println("價錢為", price)
		fmt.Println("價錢好貴")
	} else {
		fmt.Println("價錢為", price)
		fmt.Println("價錢可以接受")
	}
	// fmt.Println("價錢為", price) // 這邊會拿不到price

	n1 := 10
	n2 := 15
	if n1 >= 10 && n2 <= 20 {
		fmt.Println("hello~")
	}

	if ((n1+n2)%3 == 0) && ((n1+n2)%5 == 0) {
		fmt.Println("為3, 5公倍數")
	} else {
		fmt.Println("非3, 5公倍數")
	}

	var year int
	fmt.Println("輸入年份（西元）")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) ||
		year%400 == 0 {
		fmt.Println(year, "是閏年")
	} else {
		fmt.Println(year, "不是是閏年")
	}

	var grade byte
	fmt.Println("輸入成績")
	fmt.Scanln(&grade)
	if grade == 100 {
		fmt.Println("成績為優")
	} else if grade >= 90 {
		fmt.Println("成績為甲")
	} else if grade >= 80 {
		fmt.Println("成績為乙")
	} else {
		fmt.Println("成績不及格")
	}

	// switch
	// 每個case後可以有多個判斷，每個判斷彼此是or的關係
	// golang沒有break，case區塊執行完後即退出
	var key byte
	fmt.Println("輸入a~g")
	// fmt.Scanln(&key) // 不能使用Scanln，預設會將輸入的值用string判斷
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	// ...
	default:
		fmt.Println("不符合")
	}

	// switch [這邊可以是一個表達式，函數回傳，算式，甚至不帶] {
	// 	case [這邊的資料型態必須與switch後的資料相同]:
	// 		//
	//  case 1, 3, 5: // case 可以有多個表達式，彼此是OR關係
	//      //
	//  case 5: 此時會錯，如果case內是常數，且先前的case常數有重複，則沒意義且不能編譯
	//      //
	// 	default: 非必須
	// 		//
	// }

	var name string = "Matt"

	switch {
	case name == "Matt":
		fmt.Println("Matt")
	case name == "matt":
		fmt.Println("matt")
	}

	var temper byte = 20
	switch {
	case temper >= 28:
		fmt.Println("很熱")
	case temper >= 18 && temper < 28:
		fmt.Println("熱")
	case temper >= 15 && temper < 18:
		fmt.Println("涼")
	default:
		fmt.Println("冷")
	}

	// fallthrough
	// golang中case執行完就會跳出，若想繼續執行下一個case區塊的話在區塊最後寫上fallthrough
	// 加上fallthrough的話下一區塊的case不會判斷，而會直接執行區塊內行為
	switch value := 9; { // switch中也可以宣告變數，不推薦
	case value%3 == 0:
		fmt.Println("可被3整除")
		fallthrough
	case value%5 == 0:
		fmt.Println("可被5整除") // 這行會被輸出，儘管9不能被5整除
	default:
		fmt.Println("非3或5的倍數")
	}

	var x interface{}
	var y = 10.0
	x = y
	switch x.(type) {
	case nil:
		fmt.Printf("x的類型 %T\n", x)
	case int:
		fmt.Println("x是int")
	case float64:
		fmt.Println("x是float64")
	case func(int) float64:
		fmt.Println("x是func(int)")
	case bool, string:
		fmt.Println("x是bool或string")
	default:
		fmt.Println("x是未知型態")
	}
}
