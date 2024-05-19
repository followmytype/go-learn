package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字串常用函數
func main() {
	var str string
	str = "hello"
	// 1. 字串長度，len(v)
	fmt.Println("str len is:", len(str)) // str len is: 5
	// golang內建編碼為utf-8，ascii字元（英文字母、數字）都佔一個字節，中文字佔有三個
	str = "hello中"
	fmt.Println("str len is:", len(str)) // str len is: 8

	// 2. 字串遍歷，如果字串內有中文，需先轉換成rune切片
	runestr := []rune(str)
	for i, v := range runestr {
		fmt.Printf("key: %d, value: %c\n", i, v)
	}

	// 3. 字串轉整數
	number, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("轉成功，", number)
	} else {
		fmt.Println("轉失敗")
	}

	// 4. 整數轉字串，轉字串最簡單，所以不會有error
	str = strconv.Itoa(1235)
	fmt.Println("整數轉字串", str)

	// 5. 字串轉byte
	bytestr := []byte("hello world")
	fmt.Printf("hello world byte is %v\n", bytestr)

	// 6. byte轉字串
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	// 7. 將數字轉成其他進制的顯示
	str = strconv.FormatInt(8, 2) // (被轉的整數，要轉的進制單位)
	fmt.Println(str)

	// 8. 字串內查詢子字串
	a := strings.Contains("seafood", "oo")
	fmt.Println(a)

	// 9. 字串中有幾個子字串
	b := strings.Count("aabbcceedddeessaaccssa", "dd")
	fmt.Println(b)

	// 10. 不區分大小寫的字串比較（原先==是有區分大小寫的
	a = strings.EqualFold("aBc", "AbC")
	fmt.Println(a)

	// 11. 查詢子字串，返回第一個符合的位置，未找到的回傳-1
	c := strings.Index("seafoodoofafjieoosadfodfao", "oo")
	fmt.Println(c)

	// 12. 查詢子字串，返回最後符合的位置，未找到的回傳-1
	c = strings.LastIndex("seafoodoofafjieoosadfodfao", "oo")
	fmt.Println(c)

	// 13. 替換字串內指定字串(目標字串，要被替換字串，替換字串，替換數量(-1代表全部))
	// 會返回新的字串，原先處理的字串並不會被改變
	fmt.Println(strings.Replace("go go go alalala", "go", "hey", 1))

	// 14. 切割字串，根據給定的字串
	fmt.Println(strings.Split("azzbzzczzdzzezz", "zz")) // [a, b, c, d, e, '']

	// 15. 大小寫轉換
	fmt.Println(strings.ToLower("GAagAAGoisGJIOJ"), strings.ToUpper("faEFJDdsfAdffow"))
}
