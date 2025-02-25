package main

import (
	"flag"
	"fmt"
	"os"
)

// 獲取命令列參數
func main() {
	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Printf("第%d個參數: %v\n", i, v)
	}


	var user string
	var pwd string
	var host string
	var port int

	// 接收命令列參數 -u 後面的值
	// 第三個參數代表預設
	// 第四個代表說明
	flag.StringVar(&user, "u", "", "用戶名，預設為空")
	flag.StringVar(&pwd, "p", "", "密碼")
	flag.StringVar(&host, "h", "localhost", "主機名稱")
	flag.IntVar(&port, "port", 3306, "端口")

	// 上面都還只是定義命令參數，需呼叫這個方法才能正式將值正式設到變數中
	flag.Parse()

	fmt.Println(user)
}
