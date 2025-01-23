package main

import (
	"encoding/json"
	"fmt"
)

// struct 結構體
// struct 本身是值類型
// 結構體內的屬性可為 字串、整數、布林等等
// 也能為slice, map, point，但須注意這些預設都會是空值nil
type Cat struct {
	Name string
	Age int
	slice []int
	ptr *int
	mapp map[string]string
}

type Person struct {
	Name string
	Age int
}

func main() {
	var c1 Cat
	fmt.Println(c1)
	if c1.slice == nil {
		fmt.Println("slice of struce is nil")
	}

	if c1.ptr == nil {
		fmt.Println("ptr of struce is nil")
	}

	// c1.slice[0] = 100 // error
	// 需要先賦予空間
	c1.slice = make([]int, 5)
	c1.slice[0] = 100 // ok
	/*
	c1.slice = make([]int, 0)
	c1.slice[0] = 100 // error, 因為初始長度不夠
	c1.slice = append(c1.slice, 123) // 需要這樣寫
	*/
	fmt.Println("slice is:", c1.slice)

	if c1.mapp == nil {
		fmt.Println("map of struce is nil")
	}
	// c1.mapp["aa"] = "AA" // error, 因為沒有分配空間
	c1.mapp = make(map[string]string) // 長度無所謂，他會自己增長
	c1.mapp["aa"] = "AA" // ok
	fmt.Println(c1.mapp)

	// struct是值類型，所以不會影響到原先的變數
	p1 := Person{Name: "aa", Age: 12}
	p2 := p1
	p2.Name = "bb"
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)


	// 宣告的幾種方式
	// var pp Person
	// var pp Person = Person{}
	// var pp = Person{Name: "AAA"}
	// pp := Person{}

	// 這會是一個指針
	// var pptr *Person = new(Person) // *Person可以省略
	pptr := new(Person)
	// 因為他是指針型態，所以標準的賦值寫法，要先取值
	(*pptr).Age = 123
	// 但因為編寫方便，所以可以這樣寫，底層會自動處理
	pptr.Name = "AAA"
	fmt.Println(pptr) // 輸出前會多一個&，代表是指針
	fmt.Println(*pptr)

	// 結構體轉換，需擁有相同的屬性，名稱、類型都要一樣才能轉換
	AA := A{Param: 123}
	BB := B(AA)
	CC := C(BB)
	fmt.Printf("AA value: %v, type: %T\n", AA, AA)
	fmt.Printf("BB value: %v, type: %T\n", BB, BB)
	fmt.Printf("CC value: %v, type: %T\n", CC, CC)

	monster := Monster{Name: "魔王", Color: "紅色", Blood: 100}
	monsterJson, _ := json.Marshal(monster)
	fmt.Println(string(monsterJson))
	monster.sayName()
	monster.deductBlood() // 等於(*monster).deductBlood(), 但因為編寫方便，所以可以省略
	fmt.Println(monster)
}

type A struct {
	Param int
}
type B struct {
	Param int
}
// 新定義一個結構體C，內部構造依然是A
type C A

type Monster struct {
	Name string `json:"name"`
	Color string `json:"color"`
	Blood int `json:"blood"`
}

// 屬於monster結構體的方法，會將結構體本身複製一份當作參數使用
func (m Monster) sayName() {
	fmt.Println("我是", m.Name)
}

// 屬於monster結構體的方法，會將結構體本身帶入使用，當更改結構體內屬性值時，會一併改變外部
func (m *Monster) deductBlood() {
	m.Blood -= 10
}

// 如果一個類型實現了 String() string 方法，則在print出來時會自動套用，與php __toString()一樣
func (m Monster) String() string {
	return fmt.Sprintf("Monster Name: %v, Color: %v, Blood: %d", m.Name, m.Color, m.Blood)
}
