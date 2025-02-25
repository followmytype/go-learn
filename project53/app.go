package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var Arr = make(map[int]int, 10)
var lock sync.Mutex

// goroutine
// 進程：進程是計算機中程序的一個實例，它包含程序的代碼和執行程序所需的所有資源（例如記憶體、文件描述符等
// 	    比如打開一個vscode就是一個進程，可同時打開多個vscode，彼此之間獨立
// 線程：線程是進程中的一個執行單位。它共享進程的地址空間和資源，但擁有獨立的執行路徑。
//      vscode內有許多功能，功能可同時進行，就是一個線程
// 協程：協程是一種比線程更輕量級的並發單位。它協作執行，主動讓出控制權給其他協程，而不是被搶占調度。
//      有獨立的棧空間、共享程式堆空間、調度用戶控制、輕量級線程

// 併發：看起來像是同時執行，但實際上在特定時刻只有一個任務在執行，通過任務切換來實現多任務處理。
// 並行：真正的同時執行多個任務，所有任務在同一時刻都在運行，依賴硬件支持。
func main() {
	// 這種寫法會在test輸出完後才會輸出main
	// test("test")
	// test("main")
	// --------------

	// 需要讓test輸出與main輸出同時進行
	go test("test") // 需要背景執行的行為前面放個go，就能建立一個協程，讓他進入到背景執行
	test("main")
	// goroutine中會隨著主線程結束而一起消滅

	//------
	// 獲取可操作的cpu個數
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu 個數:", cpuNum)
	// 可設定cpu最大使用個數
	// runtime.GOMAXPROCS(cpuNum-1)

	// --------
	// 以下寫法可能會出現錯誤
	// fatal error: concurrent map writes
	// 原因是開啟的多個協程有機會同時對Arr進行寫入操作，進而導致併發問題
		// for i:=1;i<=200;i++ {
		// 	go setArrValue(i)
		// }

		// time.Sleep(10*time.Second)
		// for i, v := range Arr {
		// 	fmt.Println(i, ":", v)
		// }
	
	// 解決方法是利用鎖; 11行
	// 在
	for i:=1;i<=50;i++ {
		go setArrValueWithLock(i)
	}

	time.Sleep(5*time.Second)
	// 這裡加鎖是因為上面我們自己認知五秒就能做完所有協程，但這是人為的主觀認定，對程式來說並不是
	// 還是有可能出現資源競爭的問題，所以要加鎖
	lock.Lock()
	for i, v := range Arr {
		fmt.Println(i, ":", v)
	}
	lock.Unlock()
}

func test(input string) {
	for i := 0; i < 3; i++ {
		println(i, "hello", input)
		time.Sleep(1 * time.Second)
	}
}

func setArrValue(n int) {
	result := 1
	for i:=1;i<=n;i++ {
		result *= i
	}
	Arr[n] = result
}

func setArrValueWithLock(n int) {
	result := 1
	for i:=1;i<=n;i++ {
		result *= i
	}
	lock.Lock()
	Arr[n] = result
	lock.Unlock()
}
