package main

import (
	"fmt"
	"reflect"
)

type Aint int

type people struct {
	age int
	name string
}

type Ypeople people
type Wpeople people

func (p people) eat() {
	fmt.Println("吃了")
}

func (p people) talk() {
	fmt.Println("说话")
}

func (p Ypeople) eat() {
	fmt.Println("Ypeople 吃了")
}

func (p Wpeople) eat() {
	fmt.Println("Wpeople eat")
}

type NewInterface interface{
	talk()
}

type bird struct {
	age int
}

func (b bird) talk() {
	fmt.Println("gu gu gu", b.age)
}

func say(sb NewInterface) {
	sb.talk()
	switch sb.(type) {
case people:
	fmt.Println("他是人")
case bird:
	fmt.Println("它是鸟")
case *bird:
	fmt.Println("它是鸟*")
default:
	fmt.Println("Unknown", reflect.TypeOf(sb))
}
}

func main() {
	fmt.Println("Course 4")

	var v_Aint Aint
	fmt.Println(v_Aint, reflect.TypeOf(v_Aint))//0 main.Aint

	var xiaoli people
	xiaoli.eat()

	var xiaoY Ypeople
	//Ypeople.eat()//Ypeople.eat undefined (type Ypeople has no method eat)
	xiaoY.eat()

	var rob Wpeople
	rob.eat()


	//Interface
	fmt.Println("-----------Interface")
	var xiaoHuang = new(Ypeople)
	xiaoHuang.name = "xiao Huang"
	xiaoHuang.age = 6
	fmt.Println(xiaoHuang)//&{6 xiao Huang}

	var t_intera interface{}
	t_intera = xiaoHuang//接口可以付给任何类型
	fmt.Println(t_intera, reflect.TypeOf(t_intera))//&{6 xiao Huang} *main.Ypeople
	var t_int int = 8
	t_intera = t_int
	fmt.Println(t_intera, reflect.TypeOf(t_intera))//8 int

	var v_nl NewInterface
	//var v_int int = 9
	//v_nl = v_int
	//fmt.Println(v_nl, reflect.TypeOf(v_nl))//9 int

	var p = new(people)
	v_nl = p
	fmt.Println(v_nl, reflect.TypeOf(v_nl))//&{0 } *main.people

	
	//var gezhi *bird//runtime error: invalid memory address or nil pointer dereference
	//var gezhi = new(bird)//OK
	var gezhi bird//OK
	gezhi.age = 3
	v_nl = gezhi
	fmt.Println(v_nl, reflect.TypeOf(v_nl))//&{3} *main.bird
	v_nl.talk()//gu gu gu 3

	say(gezhi)//gu gu gu 3
	say(xiaoli)
}