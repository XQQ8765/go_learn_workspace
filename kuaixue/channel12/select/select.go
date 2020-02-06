package main

import "fmt"
import "time"

//《快学 Go 语言》第 12 课 —— 通道 https://zhuanlan.zhihu.com/p/51710515
/*
select 语句，它可以同时管理多个通道读写，
如果所有通道都不能读写，它就整体阻塞，只要有一个通道可以读写，它就会继续。
*/
func send(ch chan int, gap time.Duration) {
 i := 0
 for {
  i++
  ch <- i
  time.Sleep(gap)
 }
}

func recv(ch1 chan int, ch2 chan int) {
 for {
  select {
   case v := <- ch1:
    fmt.Printf("recv %d from ch1\n", v)
   case v := <- ch2:
	fmt.Printf("recv %d from ch2\n", v)
	//非阻塞读写需要依靠 select 语句的 default 分支,
	//当通道空时，读操作不会阻塞；当通道满时，写操作也不会阻塞
   default:
  }
 }
}

func main() {
 var ch1 = make(chan int)
 var ch2 = make(chan int)
 go send(ch1, time.Second)
 go send(ch2, 2 * time.Second)
 recv(ch1, ch2)
}