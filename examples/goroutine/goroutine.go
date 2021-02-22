package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		function example
	*/

	// let's start

	//////////////////////////////////////////////////
	// goroutine statement

	// syncronized function call
	MyFunc1("SyncFunctionCall")

	// asyncronized function call
	go MyFunc1("AsyncFunctionCall_1") // excute goroutine, then call MyFunc1
	go MyFunc1("AsyncFunctionCall_2") // excute goroutine, then call MyFunc1
	go MyFunc1("AsyncFunctionCall_3") // excute goroutine, then call MyFunc1

	// wait 3 second
	time.Sleep(time.Second * 3)

	fmt.Printf("goroutine next step  \n")

	//////////////////////////////////////////////////
	// anonymous function goroutine statement
	// create WaitGroup 생성
	var wait sync.WaitGroup
	wait.Add(2) // wait for 2 waitgroup, notice!!! 2 == need two wait.Done()

	// anonymous function goroutine
	go func() {
		defer wait.Done() //when function finish, call Done()
		fmt.Println("anonymous function 001 called")
	}() // excute goroutine, then call anonymous function

	// anonymous function goroutine with parameter
	go func(strParam string) {
		defer wait.Done() //when function finish, call Done()
		fmt.Println("anonymous function 002 called, param:", strParam)
	}("helloworld") // excute goroutine, then call this anymous funciton with "helloworld"

	wait.Wait() //wait goroutines ( above go rountines)

	fmt.Printf("goroutine done  \n")
}

//////////////////////////////////////////////////
// function statement - base
func MyFunc1(strParam string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("MyFunc4 ....., index(%d) strParam:(%s) \n", i, strParam)
	}
}
