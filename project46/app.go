package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打開既有文件並追加內容
func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0666) // 這個文件存在, O_APPEND會增加文件內容
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return
	}
	defer file.Close()

	str := "append content\n"
	// 使用帶緩存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因為wirter是帶緩存的，實際上要寫入的字串還在Writer.buf裡
	// 必須得Flush()才能確實寫入
	writer.Flush()
}
