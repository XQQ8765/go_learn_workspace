package main

import (
	"fmt"
)

//《快学 Go 语言》第 12 课 —— 通道 https://zhuanlan.zhihu.com/p/51710515
/*
通道写安全 - 向一个已经关闭的通道执行写操作会抛出异常
*/
func send(ch chan int) {
	i := 0
	for {
	 i++
	 ch <- i//panic: send on closed channel
	}
   }
   
   func recv(ch chan int) {
	value := <- ch
	fmt.Println(value)
	value = <- ch
	fmt.Println(value)
	close(ch)
   }
   
   func main() {
	var ch = make(chan int, 4)
	go recv(ch)
	send(ch)
   }