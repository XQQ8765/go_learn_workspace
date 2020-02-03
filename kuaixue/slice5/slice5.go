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
}