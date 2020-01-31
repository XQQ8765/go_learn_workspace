package main

import (
	"fmt"
	"net"
	//"log"
)

func main() {
	var protocal = "tcp"
	fmt.Println("Start tcp.")

	var addrstr = "127.0.0.1:8081"
	//var addrstr = ":#0"
	var laddr, err = net.ResolveTCPAddr(protocal, addrstr)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println("Resolve succ.")

	var listar, liserr = net.ListenTCP(protocal, laddr)
	if (liserr != nil) {
		fmt.Println("liserr:", liserr)
		return
	}
	fmt.Println("ListenTCP succ.")

	for {
		conn, err := listar.Accept()
		if (err != nil) {
			fmt.Println(err)
			return
		}
		go Read(conn)
	}
}

func Read(conn net.Conn) {
	for {
		var bytes = make([]byte, 1024)
		n, err := conn.Read(bytes)
		if (err != nil) {
			fmt.Println(err)
			conn.Close()
			return
		}
		//log.Println(string(bytes[:n]))
		fmt.Println("Recevied msg:", string(bytes[:n]))
		var writeerr = Write(conn, bytes[:n])
		if (writeerr != nil) {
			return
		}
	}
}

func Write(conn net.Conn, bytes []byte) error {
	_, err := conn.Write(bytes)
	if (err != nil) {
		fmt.Println("Error in Write:", err)
		conn.Close()
		return err
	}
	return nil
}
