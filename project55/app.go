package main

import (
	"fmt"
)

// 有數據寫入：讀取操作會等待，直到有數據被寫入channel。一旦有數據寫入，讀取操作就會立即獲取該數據並繼續執行。
// channel被關閉：如果channel被關閉（close(ch)），讀取操作會繼續讀取channel中已經存在的數據。一旦所有數據都被讀取完，讀取操作將返回該channel類型的零值，並且不再阻塞。
// 無數據且未關閉：如果channel中沒有數據，且channel也沒有被關閉，那麼讀取操作會一直阻塞等待，直到有數據被寫入或channel被關閉。這種情況下，讀取操作的等待時間理論上是無限的，直到有新的數據寫入或channel被關閉。
func writeData(intChan chan int) {
	for i:=0;i<100;i++{
		intChan <- i
		fmt.Println("write data:", i)
	}
	// 如果channel沒關閉，讀取的人會一直等待新值輸入
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		// 如果intChan內沒有值，也未被close，操作將會被阻塞，直到獲得數據
		// 如果intChan內無值且已經關閉，則操作會立即執行，ok為false
		v, ok := <- intChan 
		if !ok {
			break
		}
		fmt.Println("get data:", v)
	}
	exitChan <- true
	close(exitChan)
}

// channel 練習
func main() {
	intChan := make(chan int, 100)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	fmt.Println("waiting read and write")
	for range exitChan {
		fmt.Println("write and read data complete")
	}
}
