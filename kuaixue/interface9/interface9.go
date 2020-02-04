package main

import "fmt"

//《快学 Go 语言》第 9 课 —— 接口 https://zhuanlan.zhihu.com/p/50942676
func main() {
	fmt.Println("interface9 course.")
	interfaceExam0()

	emptyInterface()

	interfaceDynamic()

	interfaceAssign()

	interfacePointAssign()
}


// 可以闻
type Smellable interface {
	smell()
  }
  
  // 可以吃
  type Eatable interface {
	eat()
  }
  
  // 苹果既可能闻又能吃
  type Apple struct {}
  
  func (a Apple) smell() {
	fmt.Println("apple can smell")
  }
  
  func (a Apple) eat() {
	fmt.Println("apple can eat")
  }
  
  // 花只可以闻
  type Flower struct {}
  
  func (f Flower) smell() {
	fmt.Println("flower can smell")
  }

func interfaceExam0() {
	fmt.Println("-------------interface样例0：interfaceExam0()-------------")
	/*
	Go 语言的接口是隐式的，只要结构体上定义的方法在形式上（名称、参数和返回值）和接口定义的一样，
	那么这个结构体就自动实现了这个接口，我们就可以使用这个接口变量来指向这个结构体对象。
	*/
	var s1 Smellable
	var s2 Eatable
	var apple = Apple{}
	var flower = Flower{}
	s1 = apple
	s1.smell()
	s1 = flower
	s1.smell()
	s2 = apple
	s2.eat()
}


/*
如果一个接口里面没有定义任何方法，那么它就是空接口，任意结构体都隐式地实现了空接口。
Go 语言为了避免用户重复定义很多空接口，它自己内置了一个，这个空接口的名字特别奇怪，叫 interface{} 

空接口里面没有方法，所以它也不具有任何能力，其作用相当于 Java 的 Object 类型，可以容纳任意对象，它是一个万能容器。
*/
func emptyInterface() {
	fmt.Println("-------------空接口：emptyInterface()-------------")
	//一个字典的 key 是字符串，但是希望 value 可以容纳任意类型的对象，类似于 Java 语言的 Map 类型，
	//这时候就可以使用空接口类型 interface{}
	// 连续两个大括号，是不是看起来很别扭
    var user = map[string]interface{} {
        "age": 30,
        "address": "Beijing Tongzhou",
        "married": true,
    }
    fmt.Println(user)
    // 类型转换语法来了
    var age = user["age"].(int)
    var address = user["address"].(string)
    var married = user["married"].(bool)
    fmt.Println(age, address, married)
}


/*
用接口来模拟多态
*/
type Fruitable interface {
    eat()
}

type Fruit struct {
    Name string  // 属性变量
    Fruitable  // 匿名内嵌接口变量
}

func (f Fruit) want() {
    fmt.Printf("I like ")
    f.eat() // 外结构体会自动继承匿名内嵌变量的方法
}

type Apple2 struct {}

func (a Apple2) eat() {
    fmt.Println("eating apple")
}

type Banana struct {}

func (b Banana) eat() {
    fmt.Println("eating banana")
}

func interfaceDynamic() {
	fmt.Println("-------------接口模拟多态：interfaceDynamic()-------------")
    var f1 = Fruit{"Apple", Apple2{}}
    var f2 = Fruit{"Banana", Banana{}}
    f1.want()
    f2.want()
}


type Rect struct {
    Width int
    Height int
}
func interfaceAssign() {
	fmt.Println("-------------接口赋值：interfaceAssign()-------------")
	/*
	变量赋值本质上是一次内存浅拷贝，切片的赋值是拷贝了切片头，
	字符串的赋值是拷贝了字符串的头部，而数组的赋值呢是直接拷贝整个数组
	*/
	var a interface {}
    var r = Rect{50, 50}
    a = r

    var rx = a.(Rect)
    r.Width = 100
    r.Height = 100
	fmt.Println(rx)//{50, 50}//不变,数据内存的复制 —— 浅拷贝
}

func interfacePointAssign() {
	fmt.Println("-------------接口指针赋值：interfaceAssign()-------------")
	var a interface {}
    var r = Rect{50, 50}
    a = &r // 指向了结构体指针

    var rx = a.(*Rect) // 转换成指针类型
    r.Width = 100
    r.Height = 100
    fmt.Println(rx)//&{100 100}
}
