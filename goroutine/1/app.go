package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var Arr = make(map[int]int, 10)
var lock sync.Mutex

func setArrValue(n int) {
	result := 1
	for i:=1;i<=n;i++ {
		result *= i
	}
	// 鎖
	lock.Lock()
	Arr[n] = result
	// 解鎖 
	lock.Unlock()
}

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu 個數:", cpuNum)
	// 設定cpu使用個數
	// runtime.GOMAXPROCS(cpuNum-1)

	// 以下寫法可能會出現錯誤
	// fatal error: concurrent map writes
	// 原因是開啟的多個協程有機會同時對Arr進行寫入操作，進而導致併發問題
	for i:=1;i<=200;i++ {
		go setArrValue(i)
	}

	time.Sleep(10*time.Second)
	for i, v := range Arr {
		fmt.Println(i, ":", v)
	}
	// ---

	// 解決方法一，使用鎖 sync
	lock.Lock()
	for i, v := range Arr {
		fmt.Println(i, ":", v)
	}
	lock.Unlock()
}
