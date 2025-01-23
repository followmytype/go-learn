package main

import (
	"fmt"
)

// 二元搜尋
func main() {
	temp := []int{12,43,345,2233, 2248, 9898}
	target := 44
	binarySearch(temp, 0, len(temp)-1, target)
}

func binarySearch(arr []int, left, right int, target int) {
	if left > right {
		fmt.Println("no")
		return
	}

	middle := (left+right)/2
	if arr[middle] > target {
		binarySearch(arr, left, middle-1, target)
	} else if arr[middle] < target {
		binarySearch(arr, middle+1, right, target)
	} else {
		fmt.Println("success, index is:", middle)
	}
}
