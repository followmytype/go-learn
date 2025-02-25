package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Aaaa string
	Bbbb int
}

func structJson() string {
	test := Test{
		Aaaa: "hello",
		Bbbb: 123,
	}
	json, err := json.Marshal(test)
	if err != nil {
		fmt.Printf("json marshal fail: %v", err)
		return ""
	}
	fmt.Println(string(json))

	return string(json)
}

func mapJson() string {
	test := map[string]interface{}{
		"aaa": "bafdaf",
		"fdss": 123,
		"23": []string{"fa", "ioo"},
	}
	json, err := json.Marshal(test)
	if err != nil {
		fmt.Printf("json marshal fail: %v", err)
		return ""
	}
	fmt.Println(string(json))
	return string(json)
}

func sliceJson() {
	test := []int{1,2,3,4}
	json, err := json.Marshal(test)
	if err != nil {
		fmt.Printf("json marshal fail: %v", err)
		return
	}
	fmt.Println(string(json))
}

// json序列化
func main() {
	structJsonStr := structJson()
	mapJsonStr := mapJson()
	sliceJson()

	var TestStruct Test
	err := json.Unmarshal([]byte(structJsonStr), &TestStruct)
	if err != nil {
		fmt.Printf("json 反序列失敗: %v", err)
		return
	}
	fmt.Println(TestStruct)

	// 這個map不需要make，因為在unmarshal內會自動做make
	var testMap map[string]interface{}
	err = json.Unmarshal([]byte(mapJsonStr), &testMap)
	if err != nil {
		fmt.Printf("json 反序列失敗: %v", err)
		return
	}
	fmt.Println(testMap)

}
