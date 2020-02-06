package main

import (
	"fmt"
	"math/rand"
	"time"
	//"runtime"
	//"os"
)

//《快学 Go 语言》第 12 课 —— 通道 https://zhuanlan.zhihu.com/p/51710515
func main() {
	channel1()

	//channelSendRecv()

	closeChannel1()

	closeChannel2()
}

func channel1() {
	fmt.Println("-------------channel1()-------------")
	/*
	通道作为容器，它可以像切片一样，使用 cap() 和 len() 全局函数获得通道的容量和当前内部的元素个数。
	通道一般作为不同的协程交流的媒介，在同一个协程里它也是可以使用的。
	*/
	var ch chan int = make(chan int, 4)
	for i:=0; i<cap(ch); i++ {
		ch <- i   // 写通道
	}
	for len(ch) > 0 {
		var value int = <- ch  // 读通道
		fmt.Println(value)
	}
}


/*
读写阻塞
*/
func send(ch chan int) {
	for {
	 var value = rand.Intn(100)
	 ch <- value
	 fmt.Printf("send %d\n", value)
	}
   }
   
   func recv(ch chan int) {
	for {
	 value := <- ch
	 fmt.Printf("recv %d\n", value)
	 time.Sleep(time.Second)
	}
   }
   
func channelSendRecv() {
	fmt.Println("-------------channelSendRecv()-------------")
	var ch = make(chan int, 1)
	// 子协程循环读
	go recv(ch)
	// 主协程循环写
	send(ch)
}

func closeChannel1() {
	fmt.Println("-------------closeChannel1()-------------")
	/*
	Go 语言的通道有点像文件，不但支持读写操作， 还支持关闭。
	读取一个已经关闭的通道会立即返回通道类型的「零值」，而写一个已经关闭的通道会抛异常。
	如果通道里的元素是整型的，读操作是不能通过返回值来确定通道是否关闭的。
	*/
		var ch = make(chan int, 4)
		ch <- 1
		ch <- 2
		close(ch)
	   
		value := <- ch
		fmt.Println(value)//1
		value = <- ch
		fmt.Println(value)//2
		value = <- ch
		fmt.Println(value)//0
}

func closeChannel2() {
	/*
	for range 遍历通道
	当通道空了，循环会暂停阻塞，当通道关闭时，阻塞停止，循环也跟着结束了。
	当循环结束时，我们就知道通道已经关闭了。
	*/
	fmt.Println("-------------closeChannel2()-------------")
	var ch = make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)
   
	// for range 遍历通道
	for value := range ch {
	 fmt.Println(value)
	}
   }