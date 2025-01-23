package main

import (
	"fmt"
)

// map是引用類型的，若作為函式參數傳遞，內部改變一併影響外部
func main() {
	test := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}
	fmt.Println("before:", test)
	changeMapValue(test)
	fmt.Println("after:", test)
}

func changeMapValue(input map[string]string) {
	input["z"] = "ZZZ"
}
