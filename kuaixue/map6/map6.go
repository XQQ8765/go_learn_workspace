package main

import "fmt"

//《快学 Go 语言》第 6 课 —— 字典 https://zhuanlan.zhihu.com/p/50047198
func main() {
	fmt.Println("map6 course.")
	//创建字典
	createMap()

	//字典读写
	mapRW()

	mapKeyExists()

	mapIter()

	mapKeysValues()
}

func createMap() {
	fmt.Println("-------------创建字典：createMap()-------------")
	//创建字典可以使用 make 函数
	//使用 make 函数创建的字典是空的，长度为零，内部没有任何元素。
	var m map[int]string = make(map[int]string)
	//var m = make(map[int]string, 16)//通知运行时提前分配好相应的内存, 避免字典经历的多次扩容操作
	fmt.Println(m, len(m))
	
	//给字典提供初始化的元素
	var m2 map[int]string = map[int]string{
        90: "优秀",
        80: "良好",
        60: "及格",  // 注意这里逗号不可缺少，否则会报语法错误
    }
	fmt.Println(m2, len(m2))
	
	//字典变量同样支持类型推导，上面的变量定义可以简写成
	var m3 = map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格",
	   }
	fmt.Println(m3, len(m3))
}

func mapRW() {
	fmt.Println("-------------字典读写：mapRW()-------------")
	//字典可以使用中括号来读写内部元素，使用 delete 函数来删除元素
	var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }
    // 读取元素
    var score = fruits["banana"]
    fmt.Println(score)

    // 增加或修改元素
    fruits["pear"] = 3
    fmt.Println(fruits)

    // 删除元素
    delete(fruits, "pear")
	fmt.Println(fruits)
	
	//删除操作时，如果对应的 key 不存在，delete 函数会静默处理。
	//遗憾的是 delete 函数没有返回值，你无法直接得到 delete 操作是否真的删除了某个元素。
	//你需要通过长度信息或者提前尝试读取 key 对应的 value 来得知。
	delete(fruits, "xxx")

	//读操作时，如果 key 不存在，也不会抛出异常。它会返回 value 类型对应的零值。
	//如果是字符串，对应的零值是空串，如果是整数，对应的零值是 0，如果是布尔型，对应的零值是 false。
	var scoreXXX = fruits["xxx"]
    fmt.Println(scoreXXX)
}

func mapKeyExists() {
	fmt.Println("-------------mapKeyExists()-------------")
	//Note: 你不能通过返回的结果是否是零值来判断对应的 key 是否存在，因为 key 对应的 value 值可能恰好就是零值
	var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

	//字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在。
    var score, ok = fruits["durin"]
    if ok {
        fmt.Println(score)
    } else {
        fmt.Println("durin not exists")
    }

    fruits["durin"] = 0
    score, ok = fruits["durin"]
    if ok {
        fmt.Println(score)
    } else {
        fmt.Println("durin still not exists")
    }
}

func mapIter() {
	fmt.Println("-------------字典遍历：mapIter()-------------")
	var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

    for name, score := range fruits {
        fmt.Println(name, score)
    }

    for name := range fruits {
        fmt.Println(name)
    }
}

func mapKeysValues() {
	fmt.Println("-------------mapKeysValues()-------------")
	//Go 语言的字典没有提供诸于 keys() 和 values() 这样的方法，意味着如果你要获取 key 列表，就得自己循环一下
	var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

    var names = make([]string, 0, len(fruits))
    var scores = make([]int, 0, len(fruits))

    for name, score := range fruits {
        names = append(names, name)
        scores = append(scores, score)
    }

    fmt.Println(names, scores)
}