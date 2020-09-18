package main

import "fmt"

func main() {

	/*
		constants example
	*/

	// let's start

	//////////////////////////////////////////////////
	// constant must define with assignment
	const a1 int = 100
	const a2 float32 = 102.1
	const a3 string = "103_hello_world"
	const a4 bool = true
	const a5, a6 int = 105, 106

	//////////////////////////////////////////////////
	// constant can be defined without data type
	const b1 = 201
	const b2 = 202.1
	const b3 = "203_hello_world"
	const b4 = true // false
	const b5, b6 = 205, 206

	//////////////////////////////////////////////////
	// this is simillar with enum in C++
	const (
		c1 = "301_hello_world"
		c2 = "302_hello_world"
		c3 = "303_hello_world"
	)

	// iota allocat number between 0 and over 0
	const (
		d1 = iota // 0
		d2        // 1
		d3        // 2
	)

	//////////////////////////////////////////////////
	// print
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println("\n")

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(b5)
	fmt.Println(b6)
	fmt.Println("\n")

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println("\n")

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println("\n")

}
