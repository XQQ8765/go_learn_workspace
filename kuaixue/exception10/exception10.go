package main

import (
	"fmt"
	"errors"
	"os"
)

//《快学 Go 语言》第 10 课 —— 错误与异常 https://zhuanlan.zhihu.com/p/51164928
func main() {
	createError()

	readFile()

	panicRecover()

	deferOrder()

}


type SomeError struct {
    Reason string
}

func (s SomeError) Error() string {
    return s.Reason
}

func createError() {
	fmt.Println("-------------创建error：createError()-------------")
	/*
	Go 语言规定凡是实现了错误接口的对象都是错误对象
	编写一个错误对象很简单，写一个结构体，然后挂在 Error() 方法就可以了
	*/
	//错误接口
    var err error = SomeError{"something happened"}
	fmt.Println(err)
	
	//Go 语言内置了一个通用错误类型，在 errors 包里。
	//这个包还提供了一个 New() 函数让我们方便地创建一个通用错误
	var err2 = errors.New("something happened")
	fmt.Println(err2)

	//如果你的错误字符串需要定制一些参数，可使用 fmt 包提供了 Errorf 函数
	var thing = "something"
	var err3 = fmt.Errorf("%s happened", thing)
	fmt.Println(err3)
}

func readFile() {
	/*
	defer
	*/
	fmt.Println("-------------readFile()-------------")
    // 打开文件
	var f, err = os.Open("exception10.go")
	//Go语言是以返回值的形式来通知上层逻辑来处理错误
    if err != nil {
        // 文件不存在、权限等原因
        fmt.Println("open file failed reason:" + err.Error())
        return
	}
	/*
	defer 关键字，它将文件的关闭调用推迟到当前函数的尾部执行，即使后面的代码抛出了异常，文件关闭也会确保被执行，相当于 Java 语言的 finally 语句块
	*/
    // 推迟到函数尾部调用，确保文件会关闭
    defer f.Close()
    // 存储文件内容
    var content = []byte{}
    // 临时的缓冲，按块读取，一次最多读取 100 字节
    var buf = make([]byte, 100)
    for {
        // 读文件，将读到的内容填充到缓冲
        n, err := f.Read(buf)
        if n > 0 {
            // 将读到的内容聚合起来
            content = append(content, buf[:n]...)
        }
        if err != nil {
            // 遇到流结束或者其它错误
            break
        }
    }
    // 输出文件内容
    fmt.Println(string(content))
}


var negErr = fmt.Errorf("non positive number")
// 让阶乘函数返回错误太不雅观了
// 使用 panic 会合适一些
func fact(a int) int{
    if a <= 0 {
        panic(negErr)
    }
    var r = 1
    for i :=1;i<=a;i++ {
        r *= i
    }
    return r
}
func panicRecover() {
	fmt.Println("-------------panicRecover()-------------")
	/*
	Go 语言提供了 panic 和 recover 全局函数让我们可以抛出异常、捕获异常。
	它类似于其它高级语言里常见的 throw try catch 语句，但是又很不一样，
	比如 panic 函数可以抛出来任意对象。

	Go 语言官方表态不要轻易使用 panic recover，除非你真的无法预料中间可能会发生的错误，
	或者它能非常显著地简化你的代码。简单一点说除非逼不得已，否则不要使用它。
	*/

	//recover 函数来保护它，recover 函数需要结合 defer 语句一起使用，
	//这样可以确保 recover() 逻辑在程序异常的时候也可以得到调用。
	/*
	defer func() {//捕获err,但后面的程序依然不继续运行
        if err := recover(); err != nil {
            fmt.Println("error catched", err)
        }
	}()
	*/
	defer func() {
		if err := recover(); err != nil {
			if err == negErr {
				fmt.Println("error catched", err)
			} else {
				panic(err)  // rethrow
			}
		}
	}()
	fmt.Println(fact(10))
    fmt.Println(fact(5))
    fmt.Println(fact(-5))
    fmt.Println(fact(15))
}

/*
defer 语句的执行顺序和代码编写的顺序是反过来的，也就是说最先 defer 的语句最后执行
*/
func deferOrder() {
	fmt.Println("-------------deferOrder()-------------")
	fmt.Println("open source file")
    defer func() {
        fmt.Println("close source file")
    }()

    fmt.Println("open target file")
    defer func() {
        fmt.Println("close target file")
    }()
    fmt.Println("do something here")
}