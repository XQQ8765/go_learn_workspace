package main

import (
	"fmt"
	"time"
	"runtime"
	//"os"
)

//《快学 Go 语言》第 11 课 —— 千军万马跑协程 https://zhuanlan.zhihu.com/p/51516757
func main() {
	goroutine1()

	goMaxProcs()

	numGoroutine()

	//millionRoutines()

	//routineDeadLoop()

}

func goroutine1() {
	fmt.Println("-------------创建error：goroutine1()-------------")
	//协程之间并不存在这么多的层级关系，在 Go 语言里只有一个主协程，其它都是它的子协程，子协程之间是平行关系
    fmt.Println("run in main goroutine")
    go func() {
        fmt.Println("run in child goroutine")
        go func() {
            fmt.Println("run in grand child goroutine")
            go func() {
				fmt.Println("run in grand grand child goroutine")

				/*
				在使用子协程时一定要特别注意保护好每个子协程，确保它们正常安全的运行。
				因为子协程的异常退出会将异常传播到主协程，直接会导致主协程也跟着挂掉，然后整个程序就崩溃了。
				*/
				//panic("wtf")//会导致主协程也跟着挂掉
            }()
        }()
    }()
    time.Sleep(time.Second)
    fmt.Println("main goroutine will quit")
}

func goMaxProcs() {
	fmt.Println("-------------millionRoutines()-------------")
	// 读取默认的线程数
    fmt.Println("单前线程数:", runtime.GOMAXPROCS(0))
    // 设置线程数为 10
 	//runtime.GOMAXPROCS(10)
    // 读取当前的线程数
 	//fmt.Println(runtime.GOMAXPROCS(0))
}

func numGoroutine() {
	fmt.Println("-------------numGoroutine()-------------")
    fmt.Println("单前协程数:", runtime.NumGoroutine())
    for i:=0;i<10;i++ {
        go func(){
            for {
                time.Sleep(time.Second)
            }
        }()
    }
    fmt.Println("单前协程数:", runtime.NumGoroutine())
}

func millionRoutines() {
	fmt.Println("-------------millionRoutines()-------------")
	//这个程序创建了千万个协程还没有到上限，观察内存发现占用还不到 1G，这意味着每个协程的内存占用还不到 100 字节。
		i := 1
		for {
			go func() {
				for {
					time.Sleep(time.Second)
				}
			}()
			if i % 10000 == 0 {
				fmt.Printf("%d goroutine started\n", i)
			}
			i++
		}
}

func routineDeadLoop() {
    fmt.Println("-------------func routineDeadLoop()-------------")
	//当 n 值大于 3 时，主协程将没有机会得到运行，而如果 n 值为 3、2、1，主协程依然可以每秒输出一次。
	n := 3
	//n := 4//主协程将没有机会得到运行, Why?
    for i:=0; i<n; i++ {
        go func() {
            fmt.Println("dead loop goroutine start")
            for {}  // 死循环
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("main goroutine running")
    }
}