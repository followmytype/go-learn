package main

import "fmt"

// channel使用
// channel本身就是一種對列queue
// channel是線程安全的，多個線程操作不會出現資源競爭的錯誤
// channel內容物是需要設定類型的
// channel是引用類型
// channel必須初始化才能使用，即make
func main() {
	var intChannel chan int
	intChannel = make(chan int, 3)

	fmt.Printf("intChannel value: %v\n", intChannel) // 輸出一個地址
	intChannel <- 1 // 向channel寫入數值
	intChannel <- 2
	fmt.Printf("channel len: %v, channel cap: %v\n", len(intChannel), cap(intChannel))
	intChannel <- 3
	fmt.Printf("channel len: %v, channel cap: %v\n", len(intChannel), cap(intChannel))
	// intChannel <- 222 // 這會錯誤 因為超過能存放的容量上限
	var num int
	num = <- intChannel
	fmt.Printf("channel取出的value: %v\n", num)
	num2 := <- intChannel
	num3 := <- intChannel
	fmt.Printf("num2: %v, num3: %v\n", num2, num3)
	// <- intChannel // 再取一次值會錯誤，因為channel內部已經沒有值了

	anyChannel := make(chan any, 3)
	anyChannel <- "hello"
	anyChannel <- 123
	anyChannel <- Temp{
		Name: "aaa", Age: 123,
	}
	<- anyChannel
	<- anyChannel
	newTemp := <- anyChannel
	// fmt.Println(newTemp.Name) // 此時newTemp出來的類型為any, 也就是interface{}，無法當成Temp結構體看待
	fmt.Println(newTemp.(Temp).Name)

	// 關閉channel
	// 使用close()內建函數
	// channel關閉後仍然可以取值，但不能再傳遞數值進去
	anyChannel <- "hello"
	close(anyChannel)
	// anyChannel <- "world" // 這會失敗
	fmt.Println(<- anyChannel)


	forChannel := make(chan int, 100)
	for i:=0;i<100;i++ {
		forChannel <- i
	}

	// 在走訪channel時，如果沒有關閉channel，他會出現deadlock的錯誤
	// 需要事先關閉channel
	close(forChannel)
	for v := range forChannel {
		fmt.Printf("forChannel value=%v\n", v)
	}
}

type Temp struct {
	Name string
	Age int
}
