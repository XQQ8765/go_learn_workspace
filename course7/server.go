package main

import (
	"fmt"
	"net"
	"sync"
	"strconv"
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

var AL = new(allconns)
type allconns struct {
	conns map[int]*conndata
	Count int
	Lock *sync.RWMutex
}
type conndata struct {
	conn net.Conn
	id int
	name string
}

func init() {
	AL.conns = make(map[int]*conndata)
	AL.Lock = new(sync.RWMutex)
}

//增加连接
func addConn(conn net.Conn) {
	AL.Lock.Lock()
	var conndata = new(conndata)
	conndata.conn = conn
	conndata.id = AL.Count
	conndata.name = "Conn-" + strconv.Itoa(conndata.id)
	AL.conns[AL.Count] = conndata
	AL.Count++
	AL.Lock.Unlock()
}

//删除连接
func delConn(n int) {
	AL.Lock.Lock()
	delete(AL.conns, n)
	AL.Lock.Unlock()
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
	AL.Lock.RLock()
	for k, conndata := range AL.conns {
		var conn = conndata.conn
		_, err := conn.Write(bytes)
		if (err != nil) {
			AL.Lock.RUnlock()
			delConn(k)
			AL.Lock.RLock()
			conn.Close()
			continue
		}
	}
	AL.Lock.RUnlock()
}
