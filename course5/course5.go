package main

import (
	"fmt"
	//"reflect"
	"time"
)

func say(data string) {
	fmt.Println(data)
}

func sayWithBlocked(data string, channel chan int) {
	<-channel//从通道读数据，如无数据可读，将阻塞
	fmt.Println("sayWithBlocked-", data)
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

	time.Sleep(1e9)
	fmt.Println("exit.")
}