package main

import (
	"fmt"
	"go-learn/project12/model"
	// "strconv"
)

// 標識符，意即變數、方法名稱
// 規範
// 1. 由英文字母大小寫、數字、底線 _ 組成
// 2. 數字不可開頭
// 3. 嚴格區分大小寫
// 4. 不能包含空格
// 5. 不可使用單一底線做變數名稱
// 6. 不能使用保留字
// break, default, func, interface, select,
// case, defer, go, map, struct
// chan, else, goto, package, switch
// const, fallthrough, if, range, type
// continue, for, import, return, var

// 命名規範
// package名稱盡量與目錄名稱一致，簡短有意義，不和標準庫的包名重複
// 變數名、函數名、常數名建議使用駝峰
// 包內的變數、函數、常數的訪問權限依照首字母的大小寫做區分，大寫代表public，小寫代表private
func main() {
	fmt.Println(model.ModelVar)
}

// https://medium.com/%E5%BE%AE%E5%B3%AF%E9%A3%9B%E7%BF%94/golang-go-mod-%E8%B5%B7%E6%89%8B%E5%8B%A2-39a0be969ffc
// 包的引用和go mod
// 在1.11版本之前，如果有自製的包要使用，都需要將包放在GOPATH的目錄底下（或是將包的目錄製作link到GOPATH底下）
// GOPATH
//  |- src
//     |- customPkg
//         |- file1.go
//         |- file2.go
// 這樣才能在執行檔案中引入並使用
// import customPkg

// go.mod出來後，會在專案內先使用go mod ini ProjectName
// 代表此專案的主要名稱，也是做為包引入的根目錄位置
// 然後直接在專案內新增自製的包目錄
// ProjectName
//  |- main.go
//  |- customPkg
//      |- file1.go
// 而在main.go中引入的寫法就會是
// import ProjectName/customPkg
// 因為有在先前的mod檔案中註冊了專案名稱，也代表了路徑，所以在引入時，寫上專案內的包名
