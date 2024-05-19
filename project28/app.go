package main

import (
	"fmt"
	"time"
)

// time
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 數字都是固定的
	format := "2006/01/02 15:04:05"
	fmt.Println(now.Format(format))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15 04 05"))

	fmt.Println(now.Format("01"))

	// const (
	// 	Nanosecond  Duration = 1                  納秒
	// 	Microsecond          = 1000 * Nanosecond  微秒
	// 	Millisecond          = 1000 * Microsecond 毫秒
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )

	// 如果要得到100毫秒，只能用 100 * Millisecond

	i := 0
	for {
		fmt.Println(i)
		// time.Sleep(1 * time.Second) // 每秒打印一次
		// time.Sleep(100 * time.Millisecond) // 每100毫秒打印一次
		if i == 10 {
			break
		}
		i++
	}

	fmt.Println(now.Unix())     // 時間戳，單位秒，1970/01/01至當前
	fmt.Println(now.UnixNano()) // 時間戳，單位納秒

	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}
	end := time.Now().Unix()
	fmt.Println("共經歷", end-start, "秒")
}
