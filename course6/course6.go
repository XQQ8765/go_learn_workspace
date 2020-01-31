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
	//"getstrtest"//We need to put the package "getstrtest" into folder "D:\GoPath\src"
	//. "getstrtest"//Use . can ignore the package name in the below code
	_ "getstrtest"//Only load the init() method of the package, without do any otherthings
)

//Package init method, run before main() method.
func init() {
	fmt.Println("Course 6 - init().")
}

//Can define multiple init() methods.
func init() {
	fmt.Println("Course 6 - init() -2.")
}

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

	//var handler MyHttpHandler
	//http.Handle("/", handler)
	//http.ListenAndServe(":80", nil)

	//fmt.Println(getstring())
	//fmt.Println(getstrtest.getstring())//undefined: getstrtest.getstring, because the func is private
	//fmt.Println(getstrtest.Getstring())
	//fmt.Println(Getstring())//package is deifined in ".", no need to add package name here
	//fmt.Println(getstrtest.Getstring())
}

type MyHttpHandler struct {

}
func (h MyHttpHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("你好"))
}