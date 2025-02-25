package main

import "fmt"

type Usb interface {
	Start()
	End()
	// returnValue() int
	// acceptParam(param int)
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone start usb")
}

func (p Phone) End() {
	fmt.Println("Phone end usb")
}

type Camera struct {
}

func (p Camera) Start() {
	fmt.Println("Camera start usb")
}

func (p Camera) End() {
	fmt.Println("Camera end usb")
}

type Computer struct{}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.End()
}

// interface，是一個引用類型
func main() {
	var computer Computer
	var phone Phone
	var camera Camera
	computer.Working(phone)
	computer.Working(camera)

	var c CStruct
	// 接口不能單獨創一個變數，必須得依賴其他已實現該接口的變數做初始化
	var a Ainterface = c
	var b Binterface = c
	a.Say()
	b.Hello()
}

type Ainterface interface {
	Say()
}

type Binterface interface {
	Hello()
}

// 街口可以嵌入其他接口，需將其他接口的方法一併實現
type Cinterface interface {
	Computer
	Ainterface
	Binterface
	Bye()
}

type CStruct struct {
}

func (c CStruct) Say() {println("c say()")}
func (c CStruct) Hello() {println("c hello()")}
func (c CStruct) Bye() {println("c bye()")}

// 錯誤，因為出現重複的方法名稱
type Xinter interface {
	Test01()
	Test02()
}
type Yinter interface {
	Test01()
	Test03()
}
type Zinter interface {
	Xinter
	Yinter
}
