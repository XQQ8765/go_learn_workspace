package main
import (
	"fmt"
	"reflect"
)

//结构体
type people struct {
	age int
	name string
}

//九九go语言入门教程 2
//https://v.youku.com/v_show/id_XMzI4MDA0MDc5Mg==.html?spm=a2h0j.11185381.listitem_page1.5~A
func main() {
	fmt.Println("Course 2.")
	var vInt int = 5
	fmt.Println(vInt)
	fmt.Println(vInt, reflect.TypeOf(vInt))

	var vInt32 int32 = 5
	fmt.Println(vInt32, reflect.TypeOf(vInt32))

	var vInt64 int64 = 500000001000
	fmt.Println(vInt64, reflect.TypeOf(vInt64))

	var vByte byte = 5
	//"byte" is with type "uint8", 0-255, 无负数
	fmt.Println(vByte, reflect.TypeOf(vByte))

	var vString string = "hello"
	fmt.Println(vString, reflect.TypeOf(vString))

	var vArra [3]int = [3]int{}
	vArra[0] = 3
	vArra[2] = 5
	fmt.Println(vArra, reflect.TypeOf(vArra))

	//可变数组
	var vSlice []int
	fmt.Println(vSlice, reflect.TypeOf(vSlice))
	//vSlice[0] = 1//Error: runtime error: index out of range [0] with length 0
	fmt.Println(vSlice == nil)//true, 没初始化

	var vSlice2 []int = make([]int, 2)//初始化可变数组
	fmt.Println(vSlice2, reflect.TypeOf(vSlice2))
	fmt.Println(vSlice2 == nil)//false

	var vSlice3 []int = make([]int, 2, 10)//初始化可变数组,可拓展
	vSlice3[1] = 1
	fmt.Println(vSlice3, reflect.TypeOf(vSlice3), len(vSlice3), cap(vSlice3))

	var vArraStr [3]string = [3]string{}
	vArraStr[1] = "bb"
	fmt.Println(vArraStr, reflect.TypeOf(vArraStr))

	//map
	var vMap map[int]int
	fmt.Println(vMap, reflect.TypeOf(vMap), vMap == nil)

	var vMap2 map[int]int = make(map[int]int)
	fmt.Println(vMap2, reflect.TypeOf(vMap2), vMap2 == nil)
	vMap2[2] = 6
	vMap2[3] = 9
	fmt.Println(vMap2)
	fmt.Println(vMap2[0])//0

	var vstrMap = make(map[string]int)
	vstrMap["apple"] = 5
	fmt.Println(vstrMap, reflect.TypeOf(vstrMap), vstrMap == nil)
	fmt.Println(vstrMap["apple"])

	//结构体指针类型
	var vPoint *people = new(people)
	fmt.Println(vPoint, reflect.TypeOf(vPoint), vPoint == nil)//&{0 } *main.people false
	vPoint.name = "xiaoming"
	vPoint.age = 6
	fmt.Println(vPoint, reflect.TypeOf(vPoint), vPoint == nil)//&{6 xiaoming} *main.people false

	//结构体实体类型
	var vStruct people
	fmt.Println("------------", vStruct)//&{0 } *main.people false
	vStruct.name = "wang"
	vStruct.age = 8
	fmt.Println(vStruct, reflect.TypeOf(vStruct))//{8 wang} main.people
}