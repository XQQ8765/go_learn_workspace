package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"strings"
	"os"
	//"sync"
	//"sort"
	//"math"
	//"net"
	"net/http"
)

//Course 6:常见包
func main() {
	fmt.Println("Course 6")
	var t = time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Since(t))

	strs := []string{"1w", "2w", "3w"}
	sep := ";"
	fmt.Println(strings.Join(strs, sep))

	s := strconv.Itoa(15)
	fmt.Println(s, reflect.TypeOf(s))
	var _, err = strconv.Atoi("xx")
	fmt.Println(err)
	var i1, err1 = strconv.Atoi("19")
	if (err1 == nil) {
		fmt.Println(i1, err1)
	}
	var i2, err2 = strconv.Atoi("19w")
	fmt.Println(i2, err2)//0 strconv.Atoi: parsing "19w": invalid syntax

	fmt.Println(os.Args)
	//os.Exit(2)

	var handler MyHttpHandler
	http.Handle("/", handler)
	http.ListenAndServe(":80", nil)
}

type MyHttpHandler struct {

}
func (h MyHttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("你好"))
}