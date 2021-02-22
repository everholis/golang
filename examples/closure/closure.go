package main

import (
	"fmt"
)

func main() {

	/*
		closure example
	*/

	// let's start

	var rtnum int = 0

	//////////////////////////////////////////////////
	// function statement

	// ClosureTest1 is return of MyFunc1(). in other words, ClosureTest1 is anonymous function 
	// ClosureTest1 is exist in scope of MyFunc1()() 
	ClosureTest1 := MyFunc1()

	rtnum = ClosureTest1()
	fmt.Printf("ClosureTest1 return, rtnum:(%d)\n", rtnum) // rtnum : 1
	rtnum = ClosureTest1()
	fmt.Printf("ClosureTest1 return, rtnum:(%d)\n", rtnum) // rtnum : 2
	rtnum = ClosureTest1()
	fmt.Printf("ClosureTest1 return, rtnum:(%d)\n", rtnum) // rtnum : 3

	ClosureTest2 := MyFunc1()

	rtnum = ClosureTest2()
	fmt.Printf("ClosureTest2 return, rtnum:(%d)\n", rtnum) // rtnum : 1
	rtnum = ClosureTest2()
	fmt.Printf("ClosureTest2 return, rtnum:(%d)\n", rtnum) // rtnum : 2

}

//////////////////////////////////////////////////
// function statement
// function name 	  	: MyFunc1
// function parameter 	: ()
// function return type	: func() int
// function body		:
//							{
//								i := 0
//								// return anonymous function
//								return func() int {
//									i++
//									return i
//								}
//							}

func MyFunc1() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
