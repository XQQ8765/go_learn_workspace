package main

import "fmt"

//《快学 Go 语言》第 5 课 —— 灵活的切片 https://zhuanlan.zhihu.com/p/49415852
func main() {
 //使用 make 函数创建的切片内容是「零值切片」，也就是内部数组的元素都是零值
 //var s1 []int = make([]int, 5, 8)
 //var s2 []int = make([]int, 8) // 满容切片
 var s1 = make([]int, 5, 8)
 s2 := make([]int, 8)//类型自动推导，省去类型定义以及 var 关键字
 fmt.Println(s1)
 fmt.Println(s2)


 //切片的初始化
 var s []int = []int{1,2,3,4,5}  // 满容的
 fmt.Println(s, len(s), cap(s))//函数 len() 和 cap() 可以直接获得切片的长度和容量属性


 //空切片
 //容量和长度都是零的切片，叫着「空切片」
 var ss1 []int
 var ss2 []int = []int{}
 var ss3 []int = make([]int, 0)
 fmt.Println(ss1, ss2, ss3)
 fmt.Println(len(ss1), len(ss2), len(ss3))
 fmt.Println(cap(ss1), cap(ss2), cap(ss3))
 //上面三种形式创建的切片都是「空切片」，不过在内部结构上这三种形式是有差异的，甚至第一种都不叫「空切片」
 //，而是叫着「 nil 切片」。但是在形式上它们一摸一样，用起来没有区别。
 //所以初级用户可以不必区分「空切片」和「 nil 切片」
 fmt.Println(ss1 == nil, ss2 == nil, ss3 == nil)//true false false


 //切片的赋值
 sliceAssgin()


 //切片的遍历
 sliceIter()

  //切片的追加
  sliceAppand()

  //切片切割
  sliceSplit()

  //数组变切片
  arrayToSlice()

  //copy函数复制切片
  copyFunc()

  //切片的扩容点
  sliceExpand()
}

func sliceAssgin() {
	//切片的赋值是一次浅拷贝操作，拷贝的是切片变量的三个域，你可以将切片变量看成长度为 3 的 int 型数组，数组的赋值就是浅拷贝。
 //拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	fmt.Println("-------------切片的赋值-------------")
	var s1 = make([]int, 5, 8)
 // 切片的访问和数组差不多
 for i := 0; i < len(s1); i++ {
  s1[i] = i + 1
 }
 var s2 = s1
 fmt.Println(s1, len(s1), cap(s1))//[1 2 3 4 5] 5 8
 fmt.Println(s2, len(s2), cap(s2))//[1 2 3 4 5] 5 8

 // 尝试修改切片内容
 s2[0] = 255
 fmt.Println(s1)//[255 2 3 4 5], s1[0] also changed
 fmt.Println(s2)//[255 2 3 4 5]
}

//切片的遍历切片的遍历, range
func sliceIter() {
	fmt.Println("-------------切片的遍历-------------")
	var s = []int{1,2,3,4,5}
    for index := range s {
        fmt.Println(index, s[index])
    }
    for index, value := range s {
        fmt.Println(index, value)
    }
}

//切片的追加
func sliceAppand() {
	//切片每一次追加后都会形成新的切片变量，如果底层数组没有扩容，那么追加前后的两个切片变量共享底层数组，
	//如果底层数组扩容了，那么追加前后的底层数组是分离的不共享的。
	//如果底层数组是共享的，一个切片的内容变化就会影响到另一个切片
	fmt.Println("-------------切片的追加-------------")
	var s1 = []int{1,2,3,4,5}
 fmt.Println(s1, len(s1), cap(s1))

 // 对满容的切片进行追加会分离底层数组
 var s2 = append(s1, 6)
 fmt.Println(s1, len(s1), cap(s1))
 fmt.Println(s2, len(s2), cap(s2))//[1 2 3 4 5 6] 6 10

 // 对非满容的切片进行追加会共享底层数组
 var s3 = append(s2, 7)
 fmt.Println(s2, len(s2), cap(s2))//[1 2 3 4 5 6] 6 10
 fmt.Println(s3, len(s3), cap(s3))//[1 2 3 4 5 6 7] 7 10

 //Go 编译器禁止追加了切片后不使用这个新的切片变量
}

//切片切割
func sliceSplit() {
	//切片的切割是从母切片中拷贝出一个子切片来，子切片和母切片共享底层数组
	fmt.Println("-------------切片的切割-------------")
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	var s2 = s1[:5]
	var s3 = s1[3:]
	var s4 = s1[:]
	//观察cap的不同
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))//[1 2 3 4 5] 5 7
	fmt.Println(s3, len(s3), cap(s3))//[4 5 6 7] 4 4
	fmt.Println(s4, len(s4), cap(s4))//1 2 3 4 5 6 7] 7 7

	s1[0] = 255//s2[0], s4[0]都变成255
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))//[255 2 3 4 5] 5 7
	fmt.Println(s4, len(s4), cap(s4))//255 2 3 4 5 6 7] 7 7
}

//数组变切片
func arrayToSlice() {
	//对数组进行切割可以转换成切片，切片将原数组作为内部底层数组。
	//也就是说修改了原数组会影响到新切片，对切片的修改也会影响到原数组。
	fmt.Println("-------------数组变切片-------------")
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var b = a[2:6]
    fmt.Println(b)
    a[4] = 255//注意b[2]
    fmt.Println(b)//[3 4 255 6]
}

func copyFunc() {
	fmt.Println("-------------copy函数复制切片: copyFunc-------------")
	//copy 函数，用来进行切片的深拷贝。不过其实也没那么深，只是深到底层的数组而已。
	//如果数组里面装的是指针，比如 []*int 类型，那么指针指向的内容还是共享的
	var s = make([]int, 5, 8)
	for i:=0;i<len(s);i++ {
		s[i] = i+1
	}
	fmt.Println(s)//[1 2 3 4 5]
	var d = make([]int, 2, 6)
	//copy拷贝的量是原切片和目标切片长度的较小值
	var n = copy(d, s)
	fmt.Println(n, d)//2 [1 2]

	s[0] = 255
	fmt.Println(s)//[255 2 3 4 5]
	fmt.Println(d)//[1 2]

	//指针slice的copy
	var points = make([]*int, 5, 8)
	for i:=0;i<len(s);i++ {
		var j = i+1
		points[i] = &j
	}
	fmt.Println("points", points)//points [0xc000066658 0xc000066660 0xc000066668 0xc000066670 0xc000066678]
	var dpoints = make([]*int, 2, 6)
	var dn = copy(dpoints, points)
	fmt.Println(dn, dpoints)//2 [0xc000066658 0xc000066660]
	fmt.Println("*dpoints[0]:", *dpoints[0])//1
	*points[0] = 255//指针指向的内容还是共享的
	fmt.Println("*points[0]:", *points[0])//255
	fmt.Println("*dpoints[0]:", *dpoints[0])//255
}

//切片的扩容点
func sliceExpand() {
	fmt.Println("-------------copy函数复制切片: sliceExpand-------------")
	//当比较短的切片扩容时，系统会多分配 100% 的空间，也就是说分配的数组容量是切片长度的2倍。
	//但切片长度超过1024时，扩容策略调整为多分配 25% 的空间，这是为了避免空间的过多浪费。
	s1 := make([]int, 6)
	s2 := make([]int, 1024)
	s1 = append(s1, 1)
	s2 = append(s2, 2)
	fmt.Println(len(s1), cap(s1))//7 12
	fmt.Println(len(s2), cap(s2))//1025 1280
}