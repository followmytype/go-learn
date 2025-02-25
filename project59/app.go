package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(input any) {
	rTyp := reflect.TypeOf(input) // 回傳reflect.Type類型
	fmt.Println("rTyp:", rTyp)

	rVal := reflect.ValueOf(input) // 回傳reflect.Value類型
	fmt.Println("rVal:", rVal)
	fmt.Printf("rVal value: %v, rVal type: %T\n", rVal, rVal)
	// rVal + 123 // 這會出錯要下面這麼寫
	fmt.Println(rVal.Int() + 123)

	iv := rVal.Interface()
	num := iv.(int)
	fmt.Println("input is num:", num)
}

func reflectPtr(input any) {
	val := reflect.ValueOf(input)
	fmt.Printf("val kind: %v, type: %v\n", val.Kind(), val.Type())
	// reflect包中的Elem()方法是用於獲取指向的值或者數組/切片/字典的元素類型。它通常在進行反射操作時使用來取得指針指向的具體值。
	// 簡單來說，Elem()方法返回一個reflect.Value，該值表示指向的變量所包含的值。例如，如果你有一個指向整數的指針，Elem()將返回該指針所指向的整數值。這對於操作指針類型的變量非常有用。
	val.Elem().SetInt(20)
}

func main() {
	a := 123
	reflectTest01(a)

	reflectPtr(&a)
	fmt.Println(a)

	b := "tom"
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal type: %v, kind: %v\n", rVal.Type(), rVal.Kind())
	// rVal.SetString("jack") // reflect.Value只有在值是可設置的情況下才能調用 Set 類方法
	// string屬於不可變的類型
	reflect.ValueOf(&b).Elem().SetString("jack")
	fmt.Println(b)
}
