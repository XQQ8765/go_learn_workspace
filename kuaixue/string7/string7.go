package main

import "fmt"

//《快学 Go 语言》第 7 课 —— 字符串 https://zhuanlan.zhihu.com/p/50399072
func main() {
	fmt.Println("string7 course.")

	iterByByte()

	iterByRune()

	stringInMemory()

	stringReadOnly()

	stringSplit()

	stringconv()
}

//按字节遍历
func iterByByte() {
	/*
	字符串通常有两种设计，一种是「字符」串，一种是「字节」串。
	「字符」串中的每个字都是定长的，而「字节」串中每个字是不定长的。
	Go 语言里的字符串是「字节」串，英文字符占用 1 个字节，非英文字符占多个字节。
	这意味着无法通过位置来快速定位出一个完整的字符来，而必须通过遍历的方式来逐个获取单个字符。

	Go 语言中的字符 rune 占 4 个字节.rune 类型是一个衍生类型，它在内存里面使用 int32 类型的 4 个字节存储.
	type rune int32
	*/
	fmt.Println("-------------按字节遍历：iterByByte()-------------")
	var s = "嘻哈china"
	//输出: e5 98 bb e5 93 88 63 68 69 6e 61
    for i:=0;i<len(s);i++ {
        fmt.Printf("%x ", s[i])
    }
}

//按rune字符遍历
func iterByRune() {
	fmt.Println("-------------按rune字符遍历：iterByRune()-------------")
	var s = "嘻哈china"
	//对字符串进行 range 遍历，每次迭代出两个变量 codepoint 和 runeValue。
	//codepoint 表示字符起始位置，runeValue 表示对应的 unicode 编码（类型是 rune）
	//输出: 0 22075 3 21704 6 99 7 104 8 105 9 110 10 97
    for codepoint, runeValue := range s {
        fmt.Printf("%d %d ", codepoint, int32(runeValue))
    }
}

func stringInMemory() {
	fmt.Println("-------------字节串的内存表示：stringInMemory()-------------")
	/*
	字符串的内存结构: 编译器还为它分配了头部字段来存储长度信息和指向底层字节数组的指针，结构非常类似于切片，区别是头部少了一个容量字段
	当我们将一个字符串变量赋值给另一个字符串变量时，底层的字节数组是共享的，它只是浅拷贝了头部字段
	*/
	var s1 = "hello" // 静态字面量
	var s2 = ""
	for i:=0;i<10;i++ {
	s2 += s1 // 动态构造
	}
	fmt.Println(len(s1))//5
	fmt.Println(len(s2))//50
}

func stringReadOnly() {
	fmt.Println("-------------字符串是只读的：stringReadOnly()-------------")

	/*
	字符串是只读的
	你可以使用下标来读取字符串指定位置的字节，但是你无法修改这个位置上的字节内容。
	如果你尝试使用下标赋值，编译器在语法上直接拒绝你。
	*/
	//var s = "hello"
    //s[0] = 'H'//编译不通过
}

func stringSplit() {
	fmt.Println("-------------字符串切割：stringSplit()-------------")
	/*
	字符串在内存形式上比较接近于切片，它也可以像切片一样进行切割来获取子串。子串和母串共享底层字节数组。
	*/
	var s1 = "hello world"
    var s2 = s1[3:8]
    fmt.Println(s2)
}

func stringconv() {
	//字节切片和字符串的相互转换
	fmt.Println("-------------字符串转换：stringconv()-------------")
	var s1 = "hello world"
    var b = []byte(s1)  // 字符串转字节切片
    var s2 = string(b)  // 字节切片转字符串
    fmt.Println(b)
    fmt.Println(s2)
}