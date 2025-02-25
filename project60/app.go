package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32
	Sex bool
}

func (m Monster) Print() {
	fmt.Println("start---")
	fmt.Println(m)
	fmt.Println("end---")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, score float32, sex bool) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func main() {
	aMonster := Monster{
		Name: "www",
		Age: 10,
		Score: 11.5,
		Sex: true,
	}
	testReflect(aMonster)
}

func testReflect(input any) {
	inputVal := reflect.ValueOf(input)
	inputTyp := reflect.TypeOf(input)
	if inputVal.Kind() != reflect.Struct {
		fmt.Println("not struct")
		return 
	}
	fieldNum := inputVal.NumField()
	fmt.Printf("input has %v field\n", fieldNum)
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("field %d: %v\n", i, inputVal.Field(i))
		tagVal := inputTyp.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d tag: %v\n", i, tagVal)
		}
	}

	methodNum := inputVal.NumMethod()
	fmt.Printf("input has %d method\n", methodNum)
	// method的排序是由名稱ascii碼決定
	inputVal.Method(1).Call(nil)


	var params []reflect.Value
	params = append(params, reflect.ValueOf(40))
	params = append(params, reflect.ValueOf(20))
	fmt.Println("call print method:", inputVal.Method(0).Call(params)[0].Int())
	// reflect操作幾乎都適用reflect.Value
	// 包括呼叫、傳遞參數、接收回傳值等等
}
