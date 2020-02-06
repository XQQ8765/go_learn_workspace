package main

import (
	"fmt"
)

//《快学 Go 语言》第 12 课 —— 通道 https://zhuanlan.zhihu.com/p/51710515
/*
通道写安全 - 向一个已经关闭的通道执行写操作会抛出异常, 解决单写多读的场景
确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。
*/
func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
   }
   
   func recv(ch chan int) {
	for v := range ch {
	 fmt.Println(v)
	}
   }
   
   func main() {
	var ch = make(chan int, 1)
	go send(ch)
	recv(ch)
   }