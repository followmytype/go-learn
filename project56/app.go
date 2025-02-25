package main

import (
	"fmt"
)

// 使用goroutine計算1-8000共有幾個質數
// 使用四個協程幫助計算
// 1. intChan放入待計算的數字，且在放入同時開啟計算質數的協程
// 2. intChan在數字都放入後關閉
// 3. 啟動四個協程拿取intChan內值做質數計算，因為intChan有關閉動作，所以可以放心取用直到結束
// 4. 全計算完成後，在exitChan內傳遞完成的值，注意：exitChan未關閉
// 5. 將計算結果放入resultChan，注意：resultChan未關閉
// 6. 新增一個協程監看exitChan是否已完成，若完成則代表exitChan, resultChan皆可關閉
// 6. 同時間打印質數結果
func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 2000)
	exitNum := 4
	exitChan := make(chan bool, exitNum)

	go PutNum(intChan)
	for i := 0; i < 4; i++ {
		go CheckNum(intChan, resultChan, exitChan)
	}
	go func() {
		for i := 0; i < exitNum; i++ {
			<-exitChan
		}
		close(exitChan)
		close(resultChan)
	}()

	PrintResult(resultChan)
	fmt.Println("complete")
}

func PutNum(intChan chan int) {
	for i := 2; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func CheckNum(intChan chan int, resultChan chan int, exitChan chan bool) {
	for value := range intChan {
		if isPrimeNum(value) {
			resultChan <- value
		}
	}
	exitChan <- true
}

func PrintResult(resultChan chan int) {
	total := 0
	for v := range resultChan {
		fmt.Println(v)
		total++
	}
	fmt.Println("共有:", total)
}

func isPrimeNum(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
