package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
	//"log"
)

func main() {
	fmt.Println("talk start.")

	var protocal = "tcp"
	var addrstr = "127.0.0.1:8081"
	conn, err := net.Dial(protocal, addrstr)
	if (err != nil) {
		fmt.Println("Dial error", err)
		conn.Close()
		return
	}

	for {
		_, err = conn.Write(readIO())
		if (err != nil) {
			fmt.Println("Write Error", err)
			conn.Close()
			return
		}

		var b = make([]byte, 1024)
		n, err := conn.Read(b)
		if (err != nil) {
			fmt.Println("Read Error", err)
			conn.Close()
			return
		}
		fmt.Println("Received:", string(b[0:n]))
		time.Sleep(1e9)
	}
}

var count = 0
func readIO() []byte {
	count++
	return []byte("说话," + strconv.Itoa(count)) 
}