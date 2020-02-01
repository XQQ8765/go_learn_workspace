package main

import (
	"fmt"
	"net"
	"sync"
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
		addConn(conn)
		go Read(conn)
	}
}

var conns = make(map[int]net.Conn)
var count = 0
var Lock = new(sync.RWMutex)//读写锁

//增加连接
func addConn(conn net.Conn) {
	Lock.Lock()
	conns[count] = conn
	count++
	Lock.Unlock()
}

//删除连接
func delConn(n int) {
	Lock.Lock()
	delete(conns, n)
	Lock.Unlock()
}

//读取数据
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
		Write(bytes[:n])
	}
}

//写回数据
func Write(bytes []byte) {
	Lock.RLock()
	for k, conn := range conns {
		_, err := conn.Write(bytes)
		if (err != nil) {
			Lock.RUnlock()
			delConn(k)
			Lock.RLock()
			conn.Close()
			continue
		}
	}
	Lock.RUnlock()
}
