package main

import (
)

func main() {
	// 唯寫channel
	var chan1 chan<- int
	chan1 = make(chan<- int, 1)
	chan1 <- 1

	// 唯讀channel
	var chan2 <-chan int
	chan2 = make(<-chan int, 1)
	<-chan2
}

// 通常唯讀唯寫會用在方法參數上，避免誤操
func send(chan1 chan<- int, chan2 <-chan bool) {
	// chan1 唯寫
	// chan2 唯讀
}
