package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go server()
	time.Sleep(time.Second) // 等待服務器啟動

	// 客戶端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("client: 建立服務連線")

	// 客戶端傳送資料
	message := []byte("Hello, Server!")
	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}
	fmt.Println("client: 傳送訊息", string(message))

	// 客戶端等待5秒後讀取資料
	time.Sleep(time.Second * 5)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("client: 收到訊息", string(buffer[:n]))
}

func server() {
	listener, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("server: 啟動服務器")

	for {
		fmt.Println("server: 等待連線")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			fmt.Println("建立連線")

			buffer := make([]byte, 1024)
			n, err := c.Read(buffer)
			if err != nil {
				panic(err)
			}

			fmt.Println("server: 收到訊息", string(buffer[:n]))

			// 伺服器回應資料
			response := "Hello, Client!"
			_, err = c.Write([]byte(response))
			if err != nil {
				panic(err)
			}
			fmt.Println("server: 回傳訊息", response)
			response = response + "0"
			_, err = c.Write([]byte(response))
			if err != nil {
				panic(err)
			}
			fmt.Println("server: 回傳訊息", response)
			response = response + "0"
			_, err = c.Write([]byte(response))
			if err != nil {
				panic(err)
			}
			fmt.Println("server: 回傳訊息", response)
			response = response + "0"
			_, err = c.Write([]byte(response))
			if err != nil {
				panic(err)
			}
			fmt.Println("server: 回傳訊息", response)
			response = response + "0"
			_, err = c.Write([]byte(response))
			if err != nil {
				panic(err)
			}
			fmt.Println("server: 回傳訊息", response)
		}(conn)
	}
}
