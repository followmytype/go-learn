package main

import "testing"

func TestAddUpper(t *testing.T)  {
	res := addUpper(5)
	if res != 15 {
		t.Fatalf("計算失敗，得到的值為: %v，應得的值為: %v\n", res, 55)
	}
	t.Logf("測試成功，得到的值為: %v\n", res)
} 

func TestHello(t *testing.T) {
	t.Log("hello test")
}
