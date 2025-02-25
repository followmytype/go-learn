package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 將A文件內容導入到B文件內
// 判斷文件或目錄是否存在
func main() {
	file1Path := "file1.txt"
	file2Path := "file2.txt" // 不存在會動新增

	data, err := os.ReadFile(file1Path)
	if err != nil {
		fmt.Println("open file1 error")
		return
	}
	err = os.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err: %v\n", err)
	}
}

// 自製方法判斷檔案目錄是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

// 複製一份檔案
func CopyFile(srcFilePath, targetFilePath string) (int64, error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("open source file error: %v\n", err)
		return 0, err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	targetFile, err := os.OpenFile(targetFilePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open target file error: %v\n", err)
		return 0, err
	}
	defer targetFile.Close()
	writer := bufio.NewWriter(targetFile)

	return io.Copy(writer, reader)
}
