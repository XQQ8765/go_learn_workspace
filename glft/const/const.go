// Section 02 - Lecture 13 - Chapter Review
//Reference: https://github.com/XQQ8765/glft/tree/master/sec02/lec13
package main

import (
	"errors"
	"fmt"
	//log "github.com/sirupsen/logrus"
)

const huge = 9302948902983904828902830492893482093849283402839482948290348293902942893482039482903482098990899089
const (
	/*
	ota是golang语言的常量计数器,只能在常量的表达式中使用
	See: Golang 使用 iota - https://www.jianshu.com/p/08d6a4216e96
	*/
	mon = iota
	tue = iota
	wed = iota
	thu = iota
	fri = iota
	sat = iota
	sun = iota
)

func foo(x int) {

}

func main() {
	// using defer()
	defer fmt.Println("leaving main()")
	// constants
	fmt.Println(float64(huge))
	fmt.Printf("mon: %v, tue: %v, wed: %v, thu: %v, fri: %v, sat: %v, sun: %v\n",
		mon, tue, wed, thu, fri, sat, sun)

	// complex number
	c := 11 + 4i
	fmt.Printf("real(c): %v, imag(c): %v\n", real(c), imag(c))//???

	/*function return 2 values*/
	// blank idenfitier in 'if' and 'for'
	if i, _ := div(9, 15); i > 0 {
		fmt.Println("Division resulted in non-zero quoitent")
	}

	for _, i := div(9, 15); i > 0; i-- {
		fmt.Println("Number of items remaining")
	}

	// functions
	if c, err := getItemCost(1); err == nil {
		fmt.Printf("Item 1 cost is $%.2f\n", c)
	}
	if _, err := getItemCost(0); err != nil {
		fmt.Println(err)
	}
	if _, err := getItemCost(-1); err != nil {
		fmt.Println(err)
	}

	hoo()//Note the execution order for "defer"
	
	//return value is a func
	f1 := retAfunc(2)
	f1()
	f2 := retAfunc(5)
	f2()
}

func getItemCost(itemId int) (float64, error) {
	if itemId == 0 {
		return 0.0, errors.New("Invalid parameter, itemId must be > 0")
	}
	if itemId < 0 {
		return 0.0, errors.New("Negative values not allowed, itemId must be > 0")
	}
	return 12.94, nil
}

//function return 2 values
func div(x, y int) (int, int) {
	a := x / y
	b := x % y
	return a, b
}

//Note the execution order for "defer"
func hoo() {
	defer fmt.Println("Leaving hoo()")
	defer fmt.Println("Entering hoo()")
}

//return value is a func
func retAfunc(x int) func() {
	y := 15
	z := 100
	var voo = func() {
		fmt.Println("Hello, World!")
		fmt.Printf("x = %v, y = %v\n", x, y)
	}
	fmt.Println(z)
	return voo
}