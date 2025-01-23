package main

import (
	"fmt"
	"sort"
)

// map
// map[keyType]valueType
// map是無序的
func main() {
	// 聲明是不會分配內存的，需要使用make才能進行賦值和使用
	var stringMap map[string]string
	// stringMap["hello"] = "world" // 會錯誤 panic: assignment to entry in nil map 
	// make會給空間分配
	stringMap = make(map[string]string, 10) // 第二個參數是大小，可不寫，預設為一
	stringMap["hello"] = "world"
	fmt.Println(stringMap)

	stringMap2 := make(map[string]string)
	stringMap2["aa"] = "AA"
	stringMap2["bb"] = "BB"
	stringMap2["cc"] = "CC"
	fmt.Println(stringMap2)

	stringMap3 := map[string]string{
		"xx": "XX",
		"yy": "YY",
		"zz": "ZZ",
	}
	fmt.Println(stringMap3)

	// value也能是個map
	students := map[int]map[string]string{
		1: {
			"name": "matt",
			"sex": "male",
		},
		2: {
			"name": "mina",
			"sex": "female",
		},
	}
	students[3] = map[string]string{
		"name": "tom",
		"sex": "male",
	}
	fmt.Println(students)
	// 移除指定的key
	delete(students, 2)
	fmt.Println(students)
	// 若刪除未存在的也不會出錯
	delete(students, 99)
	fmt.Println(students)

	for k, v := range students {
		fmt.Printf("%v - %v\n", k, v)
	}

	// 清空map內值：直接賦予新的map
	students = make(map[int]map[string]string)
	fmt.Println(students)

	// 取值，判斷是否有值
	val, ok := students[44]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("not exists")
	}
	// 也可以這樣寫
	if students[1] == nil {
		fmt.Println("not exists")
	} else {
		fmt.Println(students[1])
	}

	for k, v := range students {
		fmt.Printf("%v - %v\n", k, v)
	}

	// slice + map
	cities := make([]map[string]string, 0)
	cities = append(cities, map[string]string{
		"taipei": "台北",
	})
	fmt.Println(cities)

	// 因為map是無序的，所以要順序輸出的話，需要有另一個輔助變數 // go在fmt輸出中會自動排序以便於測試，但實際上他還是無序的
	// 創建另一個slice，並且將map的key放入slice中，對slice進行排序
	m := map[int]string{
        3: "three",
        1: "one",
        2: "two",
    }
	fmt.Println("這是有序的", m) // 有序的
    // 提取鍵並排序
    keys := make([]int, 0, len(m))
	fmt.Println("這是無序的")
    for key := range m {
		fmt.Println(m[key])
        keys = append(keys, key)
    }
    sort.Ints(keys)

    // 按排序後的鍵輸出
    for _, key := range keys {
        fmt.Printf("%d: %s\n", key, m[key])
    }
}
