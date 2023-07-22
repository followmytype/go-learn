建置一個go專案都需要執行init的指令
```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.20 go mod init {projectName}
```
---
利用go的docker環境幫忙編譯檔案
```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.20 go build -v
```
因為使用的docker環境為linux，編譯出來的檔案只能在linux上面執行，所以要指定為哪個作業系統編譯。以下為編譯成mac能執行的檔案：
```
docker run --rm \
-v "$PWD":/usr/src/myapp \
-w /usr/src/myapp \
-e GOOS=darwin -e GOARCH=amd64 \
golang:1.20 go build -v
```
[補充](https://ithelp.ithome.com.tw/articles/10225188)

不編譯，直接當作腳本執行
```
docker run --rm \
-v "$PWD":/usr/src/myapp \
-w /usr/src/myapp \
golang:1.20 go run {file}.go
```
---
可以使用以下指令讓檔案內的程式碼格式做調整並重新寫入
```
gofmt -w {file\path.go} 
```
