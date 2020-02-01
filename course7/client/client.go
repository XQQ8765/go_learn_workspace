package main

import (
	"fmt"
	"net"
	//"time"
	//"strconv"
	"os"
	"bufio"
	"log"
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
		go write(conn)

		var b = make([]byte, 1024)
		n, err := conn.Read(b)
		if (err != nil) {
			fmt.Println("Read Error", err)
			conn.Close()
			return
		}
		log.Println("Received:", string(b[0:n]))
	}
}

func write(conn net.Conn) {
	var reader = bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		_, err := conn.Write(bytes)
		if (err != nil) {
			fmt.Println("Write Error", err)
			conn.Close()
			return
		}
	}
}