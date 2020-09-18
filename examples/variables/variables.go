package main

import "fmt"

func UNUSED(x ...interface{}) {}

func main() {

	/*
		variables example
	*/

	// let's start

	//////////////////////////////////////////////////
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool
	var a5, a6 int

	a1 = 101
	a2 = 102.1
	a3 = "103_hello_world"
	a4 = true // false
	a5, a6 = 105, 106

	// avoid "xxxxx declared and not used" error
	UNUSED(a1)
	_ = a2
	_ = a3
	_, _, _ = a4, a5, a6

	//////////////////////////////////////////////////
	var b1 int = 201
	var b2 float32 = 202.1
	var b3 string = "203_hello_world"
	var b4 bool = true // false
	var b5, b6 int = 205, 206

	//////////////////////////////////////////////////
	var c1 = 301
	var c2 = 302.1
	var c3 = "303_hello_world"
	var c4 = true // false
	var c5, c6 = 305, 306

	//////////////////////////////////////////////////
	// assignment
	a1 = 101

	// variable definition with assignment
	d1 := 401

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
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	fmt.Println("\n")

	fmt.Println(d1)

}
