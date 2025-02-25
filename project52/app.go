package main

import ()

func addUpper(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}
