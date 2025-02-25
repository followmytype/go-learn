package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件處理
func main() {
	// os.Open這是一個簡單的函數，用於以只讀模式（read-only）打開文件。
	// 它實際上是 os.OpenFile 的一個包裝函數，等價於 os.OpenFile(name, os.O_RDONLY, 0)
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file error,", err)
	}

	fmt.Printf("file: %v\n", file) // file是一個指標變數，會輸出地址

	err = file.Close()
	if err != nil {
		fmt.Println("close file error,", err)
	}

	// 帶緩衝的方式讀取文件
	// 打開一個文件
	file2, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file error,", err)
	}
	// 記得結束前關閉
	defer file2.Close()

	// 建立一個專門查看文件的讀取器
	reader := bufio.NewReader(file2)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str) // 這邊不用Println, 因為原本文件就已經有換行符號了
	}
	fmt.Println("read file over")
	fmt.Println("")

	// 一次性打開文件
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// content是[]byte
	fmt.Println("File content:", string(content))
}

// 以下是 ioutil 中被棄用的函數及其替代方法的對照表：

// 棄用函數	替代方法	說明
// ioutil.ReadFile	os.ReadFile		讀取文件內容
// ioutil.WriteFile	os.WriteFile	寫入文件內容
// ioutil.ReadAll	io.ReadAll		讀取所有數據
// ioutil.ReadDir	os.ReadDir		讀取目錄內容
// ioutil.TempFile	os.CreateTemp	創建臨時文件
// ioutil.TempDir	os.MkdirTemp	創建臨時目錄
// ioutil.NopCloser	io.NopCloser	將 io.Reader 轉為 io.ReadCloser
// ioutil.Discard	io.Discard		丟棄數據
