package main

import (
	"bufio"
	"fmt"
	"os"
)

// 創建新文件並寫入內容
func main() {
	file, err := os.OpenFile("temp.txt", os.O_WRONLY|os.O_CREATE, 0666) // 這個文件不存在
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return
	}
	defer file.Close()

	str := "hello world\n"
	// 使用帶緩存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因為wirter是帶緩存的，實際上要寫入的字串還在Writer.buf裡
	// 必須得Flush()才能確實寫入
	writer.Flush()
}
