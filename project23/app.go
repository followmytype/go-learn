package main

import (
	"fmt"
	// init的執行順序按照包的依賴關係
	// 同一個包內可有多個init（目前依照檔案的名稱作為順序，但並沒有明確定義同包內的init執行順序，勿依賴
	initpkg "go-learn/project23/initPkg"
	// 如果只想使用該包的init，但又不會用到該包的任何變數、函示，可這樣寫
	_ "go-learn/project23/notUsed"
)

// 包內若有全局變數聲明的話，則優先執行全局變數
// 因為聲明變數會優先，而a的值來源於test()這個方法，所以會先執行到test()
var a = test()

func test() int {
	fmt.Println("test()...")
	return 123
}

// init函數，在包引入時會優先執行
// 通常作為初始化的工作
// init不可被調用
func init()  {
	fmt.Println("init()...")
}
// init很特殊，可以重複定義，執行順序由上到下
func init()  {
	println(1)
}
func init()  {
	println(2)
}
func init()  {
	println(3)
}

// init與main同在時，init優先執行
func main() {
	fmt.Println("main()...")

	fmt.Println("initPkg的Name is", initpkg.Name)
}
