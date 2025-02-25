package main

import "fmt"

func TypeJudge(input ...interface{}) {
	for i, v := range input {
		switch v.(type) {
		case float32:
			fmt.Printf("index: %v, type is float32\n", i)
		case float64:
			fmt.Printf("index: %v, type is float64\n", i)
		case int, int32:
			fmt.Printf("index: %v, type is float64\n", i)
		case bool:
			fmt.Printf("index: %v, type is bool\n", i)
		case string:
			fmt.Printf("index: %v, type is string\n", i)
		default:
			fmt.Printf("index: %v, dont know type\n", i)
		}
	}
}

func main() {
	var a float32 = 1.1
	var b float64 = 1.1
	var c = false
	var d = 234
	var e = "strrrr"
	TypeJudge(a, b, c, d, e)
}
