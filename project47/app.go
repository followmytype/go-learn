package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 打開既有文件並讀取且增加內容
func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0666) // 這個文件存在, O_RDWR代表可讀可寫
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

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
