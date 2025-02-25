package main

import (
	"fmt"

)

func test1() {
	for i:=1;i<10;i++{
		fmt.Println("test1", i)
	}
}
func test2() {
	// 需要這樣使用去捕獲錯誤
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch err", err)
		}
	} ()
	panic("test2")
}

func main() {
	go test1()
	go test2() // 如果協程內出現panic，則會影響到主線程

	ch1 := make(chan int, 10)
    ch2 := make(chan int, 5)

	for i:=0;i<10;i++ {
		ch1<-i
	}

	for i:=0;i<5;i++ {
		ch2<-i
	}

	for {
		// 使用select可以避免拿取不到deadlock的問題
		// select類似switch，會去尋找匹配成功的操作
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		default:
			fmt.Println("over")
			return
		}
	}
}
