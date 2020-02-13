package main

import (
	"fmt"
)

type AA struct {
	age int
}

//(a AA)是func的接收者，方法调用时是值拷贝，不可以改变结构体的字段值
func (a AA) SetAge(v int) {
	a.age = v
}

//(a *AA)是func的指针接收者，可以改变结构体的字段值
func (a *AA) SetAge2(v int) {
	a.age = v
}

func (a AA) GetAge() int {
	return a.age
}

func main() {
	var aa = AA{10}
	aa.SetAge(20)
	fmt.Println("aa.GetAge():", aa.GetAge())//10, 不变

	aa.SetAge2(30)
	fmt.Println("aa.GetAge():", aa.GetAge())//30
}