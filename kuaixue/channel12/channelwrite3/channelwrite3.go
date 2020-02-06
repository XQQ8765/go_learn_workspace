package main

import (
	"fmt"
	"sync"
	"time"
)

//《快学 Go 语言》第 12 课 —— 通道 https://zhuanlan.zhihu.com/p/51710515
/*
通道写安全 - 向一个已经关闭的通道执行写操作会抛出异常, 解决多写单读的场景
这个就需要使用到内置 sync 包提供的 WaitGroup 对象，它使用计数来等待指定事件完成。
*/
func send(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 计数值减一
	i := 0
	for i < 4 {
	 i++
	 ch <- i
	}
   }
   
   func recv(ch chan int) {
	for v := range ch {
	 fmt.Println(v)
	}
   }
   
   func main() {
	var ch = make(chan int, 4)
	var wg = new(sync.WaitGroup)
	wg.Add(2) // 增加计数值
	go send(ch, wg)  // 写
	go send(ch, wg)  // 写
	go recv(ch)
	// Wait() 阻塞等待所有的写通道协程结束
	// 待计数值变成零，Wait() 才会返回
	wg.Wait()
	// 关闭通道
	close(ch)
	time.Sleep(time.Second)
   }