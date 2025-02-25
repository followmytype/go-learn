package main

import (
	"fmt"
	"io"
	"net"
)

// port 0-65535
// 0為保留port
// 1-1024為固定port, 22: ssh, 23: talnet, 21: ftp, 25: smtp, 80: iis
// 1025-65535為動態port
// 伺服器盡量少開port
// 一個port只能被一個程序監聽
func main() {
	// 1. 使用tcp網路協議
	// 2. 監聽8888本地port（127.0.0.1:8888代表監聽ipv4，0.0.0.0:8888代表ipv4, ipv6都適用
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("啟動失敗: %v\n", err)
	}
	defer listener.Close()

	// 循環等待客戶端連接
	for {
		// fmt.Println("等待請求中...")
		conn, err := listener.Accept() // 等待接收客戶端連接，當沒有請求會持續阻塞在這
		if err != nil {
			fmt.Printf("連結失敗: %v\n", err)
		} else {
			fmt.Printf("連結成功: %v, 客戶端ip: %s\n", conn, conn.RemoteAddr())
			go func(conn net.Conn) {
				defer conn.Close()

				// 等待客戶端後續請求
				for {
					// 建立接收用戶傳遞資料的容器
					buf := make([]byte, 1024)
					// 如果客戶端未傳送資料，會在此阻塞等待請求，直到超時
					// 將讀取到的東西放入buf內
					// n為讀到的字節長度
					fmt.Printf("等待%s後續輸入...\n", conn.RemoteAddr())
					n, err := conn.Read(buf)
					if err != nil {
						if err == io.EOF {
							fmt.Println("用戶已退出")
						} else {
							fmt.Println("讀取失敗:", err)
						}
						return
					}
					fmt.Printf("讀取到的內容：%s\n", string(buf[:n])) // 注意：需要對buf切片做精確裁減，因為buf內可能有其他殘值
					if _, err := conn.Write([]byte("伺服器已收到"+string(buf[:n]))); err != nil {
						fmt.Println("error", err)
					}
				}
			} (conn)
		}
	}
}
