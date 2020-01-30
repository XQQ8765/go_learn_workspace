package main

import (
	"fmt"
	//"reflect"
	"time"
	"sync"
)

func say(data string) {
	fmt.Println(data)
}

func sayWithBlocked(data string, channel chan int) {
	<-channel//从通道读数据，如无数据可读，将阻塞
	fmt.Println("sayWithBlocked-", data)
}

func say0(data string, channel chan string) {
	time.Sleep(1e9)
	var text = <- channel//1
	channel <- data//2
	fmt.Println("say0 -", data, text)
}

func talk0(data string, channel chan string) {
	channel <- data//0
	a := <- channel
	fmt.Println("talk0 -", data, a)
}

var name = "aaa"
var lock = new (sync.Mutex)
func sayWithLock(data string) {
	lock.Lock()
	fmt.Println(data)
	time.Sleep(1e9)
	fmt.Println(data)
	lock.Unlock()
}
func talkWithLock(data string) {
	lock.Lock()
	fmt.Println(data)
	lock.Unlock()
}

func main() {
	fmt.Println("Course 5.")
	go say("1 Ni hao")
	go say("2 Ni hao")
	go say("3 Ni hao")
	time.Sleep(1e9)//Sleep 1s
	fmt.Println("2nd Ni hao")

	var channel = make(chan int, 0)//无缓冲通道
	go sayWithBlocked("Hello", channel)
	time.Sleep(1e9)
	channel <- 1//写入数据至通道

	var strChannel = make(chan string, 0)
	go say0("1st wai xin reng.", strChannel)
	go talk0("2st wai xin reng.", strChannel)

	go sayWithLock(name)
	lock.Lock()
	name = "bbb"
	lock.Unlock()
	go talkWithLock(name)

	time.Sleep(2e9)
	fmt.Println("exit.")
}