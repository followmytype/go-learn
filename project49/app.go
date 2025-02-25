package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 英文字元
	NumCount   int // 數字字元
	SpaceCount int // 空格字元
	OtherCount int // 其餘
}

// 統計檔案內各字元數量
func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var charCount CharCount

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, c := range str {
			switch {
			case c >= 'a' && c <= 'z':
				fallthrough
			case c >= 'A' && c <= 'Z':
				charCount.ChCount++
			case c >= '0' && c <= '9':
				charCount.NumCount++
			case c == ' ' || c == '\t':
				charCount.SpaceCount++
			default:
				// 換行或是中文字
				charCount.OtherCount++
			}
		}
	}

	fmt.Println("統計結果：", charCount)
}
