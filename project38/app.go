package main

import "fmt"

type Goods struct {
	Name  string
	Price int
}
func (g Goods) PrintName() {
	fmt.Println(g.Name)
}
type Brand struct {
	Name    string
	Address string
}
func (b Brand) PrintName() {
	fmt.Println(b.Name)
}
type TV struct {
	Goods
	Brand
}
// 也可以只給指標
type TV2 struct {
	*Goods
	*Brand
}

func main() {
	// 當一個結構體嵌套兩種結構體時，且結構體都有相同的屬性、方法名稱，在使用上都需要明確標示
	tv1 := TV{
		Goods: Goods{
			Name: "tv111",
			Price: 1100,
		},
		Brand: Brand{
			Name: "brand111",
			Address: "address 111",
		},
	}
	tv1.Goods.PrintName()
	tv1.Brand.PrintName()

	tv2 := TV2{
		Goods: &Goods{
			Name: "tv111",
			Price: 1100,
		},
		Brand: &Brand{
			Name: "brand111",
			Address: "address 111",
		},
	}
	fmt.Println(tv2) // {0xc000010180 0xc00006e020}
	tv2.Goods.PrintName() // 一樣能呼叫
	fmt.Println(tv2.Brand.Name)

	// 使用方法
	test := test{}
	test.int = 11
	fmt.Println("test int is", test.int)
}
// 可以直接用型別做匿名屬性
type test struct {
	int // 但int的匿名屬性只能出現一次，不能出現第二個int匿名屬性，否則無法判斷
}
