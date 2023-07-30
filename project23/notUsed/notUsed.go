package notused

var precomputed = [10]int{}

func init()  {
	// 常見的用法為初始化一些需要計算結果的變數內容
	var current int = 1
    precomputed[0] = current
    for i := 1; i < len(precomputed); i++ {
        precomputed[i] = precomputed[i-1] * 2
		println(precomputed[i-1])
    }

	println("hello im not used")
}
