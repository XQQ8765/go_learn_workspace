package main

import (
	"fmt"
	"unsafe"
	"math"
)

type Circle struct {
	x int
	y int
	Radius int
  }

//快学 Go 语言》第 8 课 —— 结构体 https://zhuanlan.zhihu.com/p/50654803
func main() {
	fmt.Println("struct8 course.")

	structCreate()

	structAssign()

	arrayAndSliceStruct()

	structParams()

	structMethods()

	structWithoutDynamic()
}

func structCreate() {
	fmt.Println("-------------创建结构体：structCreate()-------------")
	/*
	第一种创建形式: kv形式
	*/
	//显示指定结构体内部字段的名称和初始值来初始化结构体
	
	var c Circle = Circle {
        x: 100,
        y: 100,
        Radius: 50,  // 注意这里的逗号不能少
    }
	fmt.Printf("%+v\n", c)
	
	//那些没有指定初值的字段会自动初始化为相应类型的「零值」
	var c1 Circle = Circle {
        Radius: 50,
    }
    var c2 Circle = Circle {}
    fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	
	/*
	第二种创建形式: 顺序形式
	不指定字段名称来顺序字段初始化，需要显示提供所有字段的初值，一个都不能少。
	*/
	var c3 Circle = Circle {100, 100, 50}
	fmt.Printf("%+v\n", c3)
	
	//结构体变量和普通变量都有指针形式，使用取地址符就可以得到结构体的指针类型
	var cp *Circle = &Circle {100, 100, 50}
	fmt.Printf("%+v\n", cp)
	
	/*
	创建的第三种形式: 使用全局的 new() 函数来创建一个「零值」结构体，所有的字段都被初始化为相应类型的零值。
	*/
	var cp1 *Circle = new(Circle)//new() 函数返回的是指针类型
	fmt.Printf("%+v\n", cp1)
	
	/*
	创建的第四种形式：这种形式也是零值初始化
	*/
	var c4 Circle//c4 is not nil
	fmt.Printf("%+v\n", c4)//零值结构体是会实实在在占用内存空间的，只不过每个字段都是零值

	//nil 结构体是指结构体指针变量没有指向一个实际存在的内存。
	//这样的指针变量只会占用 1 个指针的存储空间，也就是一个机器字的内存大小。
	//var c5 *Circle = nil
}

func structAssign() {
	fmt.Println("-------------结构体赋值：structAssign()-------------")
	//结构体之间可以相互赋值，它在本质上是一次浅拷贝操作，拷贝了结构体内部的所有字段。
	var c1 Circle = Circle {Radius: 50}
    var c2 Circle = c1
    fmt.Printf("c1 %+v\n", c1)
    fmt.Printf("c2 %+v\n", c2)
    c1.Radius = 100
    fmt.Printf("c1 %+v\n", c1)//c1 {x:0 y:0 Radius:100}
    fmt.Printf("c2 %+v\n", c2)//c2 {x:0 y:0 Radius:50}//不变

	//结构体指针之间也可以相互赋值，它在本质上也是一次浅拷贝操作，不过它拷贝的仅仅是指针地址值，结构体的内容是共享的。
    var c3 *Circle = &Circle {Radius: 50}
    var c4 *Circle = c3
    fmt.Printf("c3 %+v\n", c3)
    fmt.Printf("c4 %+v\n", c4)
    c3.Radius = 100
    fmt.Printf("c3 %+v\n", c3)//c1 {x:0 y:0 Radius:100}
    fmt.Printf("c4 %+v\n", c4)//c2 {x:0 y:0 Radius:100}
}

type ArrayStruct struct {
    value [10]int
}

/*
切片头的结构体形式如下，它在 64 位机器上将会占用 24 个字节

type slice struct {
  array unsafe.Pointer  // 底层数组的地址
  len int // 长度
  cap int // 容量
}
*/
type SliceStruct struct {
    value []int
}

func arrayAndSliceStruct() {
	fmt.Println("-------------结构体中的数组和切片：structAssign()-------------")
	//数组只有「体」，切片除了「体」之外，还有「头」部。
	//切片的头部和内容体是分离的，使用指针关联起来。
    var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
    var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
    fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))//80 24
}


func expandByValue(c Circle) {
    c.Radius *= 2
}
func expandByPointer(c *Circle) {
    c.Radius *= 2
}
func structParams() {
	fmt.Println("-------------结构体的参数传递：structParams()-------------")
	//通过值传递，在函数里面修改结构体的状态不会影响到原有结构体的状态，函数内部的逻辑并没有产生任何效果。
	var c = Circle {Radius: 50}
    expandByValue(c)//{0 0 50}//不变
	fmt.Println(c)
	
	//通过指针传递就不一样。
    expandByPointer(&c)//{0 0 100}
    fmt.Println(c)
}

// 面积
func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
   }
   
   // 周长
   func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.Radius)
   }
func structMethods() {
	fmt.Println("-------------结构体方法：structMethods()-------------")
	//Go 语言不是面向对象的语言，它里面不存在类的概念，结构体正是类的替代品。
	//类可以附加很多成员方法，结构体也可以。
	var c = Circle {Radius: 50}
	fmt.Println(c.Area(), c.Circumference())
	// 指针变量调用方法形式上是一样的
	var pc = &c
	fmt.Println(pc.Area(), pc.Circumference())
   }


/*
内嵌结构体
*/
type Point struct {
    x int
    y int
}
func (p Point) show() {
  fmt.Println(p.x, p.y)
}
type CircleWithPoint struct {
    loc Point
    Radius int
}
/*
内嵌的结构体不提供名称。这时外面的结构体将直接继承内嵌结构体所有的内部字段和方法，就好像把子结构体的一切全部都揉进了父结构体一样。
匿名的结构体字段将会自动获得以结构体类型的名字命名的字段名称
*/
type CircleWithAnonPoint struct {
    Point // 匿名内嵌结构体
    Radius int
}
func nestedStrcut() {
	fmt.Println("-------------内嵌结构体：nestedStrcut()-------------")
		var c = CircleWithPoint {
			loc: Point {
				x: 100,
				y: 100,
			},
			Radius: 50,
		}
		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", c.loc)
		fmt.Printf("%d %d\n", c.loc.x, c.loc.y)
		c.loc.show()

		var c1 = CircleWithAnonPoint {
			Point: Point {
				x: 100,
				y: 100,
			},
			Radius: 50,
		}
		fmt.Printf("%+v\n", c1)
		fmt.Printf("%+v\n", c1.Point)
		fmt.Printf("%d %d\n", c1.x, c1.y) // 继承了字段
		fmt.Printf("%d %d\n", c1.Point.x, c1.Point.y)
		c1.show() // 继承了方法
		c1.Point.show()
}


/*
Go 语言的结构体没有多态性
面向对象的多态性需要通过 Go 语言的接口特性来模拟
*/
type Fruit struct {}
func (f Fruit) eat() {
    fmt.Println("eat fruit")
}

func (f Fruit) enjoy() {
    fmt.Println("smell first")
    f.eat()
    fmt.Println("clean finally")
}

type Apple struct {
    Fruit
}
func (a Apple) eat() {
    fmt.Println("eat apple")
}

type Banana struct {
    Fruit
}
func (b Banana) eat() {
    fmt.Println("eat banana")
}

func structWithoutDynamic() {
	fmt.Println("-------------内嵌结构体：structWithoutDynamic()-------------")
	var apple = Apple {}
	var banana = Banana {}
	apple.enjoy()
	banana.enjoy()
}