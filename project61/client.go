package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("連結失敗, %v\n", err)
		return
	}
	defer conn.Close()
	// stdin: 標準輸入, stdout: 標準輸出, stderr: 標準錯誤輸出
	reader := bufio.NewReader(os.Stdin) // 新增讀取器，讀取來源：標準輸入(終端)

	go PrintServer(conn)

	for {
		// 從終端讀取用戶一行字串
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("讀取文字失敗: %v\n", err)
			return
		}
		line = strings.Trim(line, "\n")
		if line == "exit" {
			break
		}
		// 將訊息發送給伺服器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Printf("發送失敗: %v\n", err)
			return
		}
		fmt.Printf("發送了%d個字節\n", n)

	}
}

func PrintServer(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}
		fmt.Println("server:", string(buf[:n]))
	}
}
