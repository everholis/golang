package main

import "fmt"

func main() {

	/*
		operator example
	*/

	// let's start

	var a1, a2, a3 int = 0, 0, 0
	var b1, b2, b3 bool = false, false, false
	var c1, c2, c3 bool = false, false, false
	var d1, d2, d3 int = 0, 0, 0
	var e1, e2, e3 int = 0, 0, 0
	var f1, f2, f3 int = 0, 0, 0

	_, _, _ = b1, b2, b3
	_, _, _ = c1, c2, c3
	_, _, _ = d1, d2, d3
	_, _, _ = e1, e2, e3
	_, _, _ = f1, f2, f3

	//////////////////////////////////////////////////
	// Arithmetic operator
	a1 = 3 + 2 - 1
	a2 = (3 * 2) / 2
	a3++

	//////////////////////////////////////////////////
	// Relational operator
	b1 = (1 != 2) // !=, < , <=, >, >=

	//////////////////////////////////////////////////
	// Logical operator
	c1 = true && false
	c2 = true || false
	c3 = !(true && false)

	//////////////////////////////////////////////////
	// Bitwise operator
	d1 = (3 & 2)  //
	d1 = 400 << 5 //bit shift operator

	//////////////////////////////////////////////////
	// Assignment operator
	e1 = 101
	e2 *= 10 // a2 = a2 * 10   // += , -= , /= , %= , &= , |= , ^= , &^= , <<= , >>=

	//////////////////////////////////////////////////
	// Pointer operator
	// nil != 0 // nil is not 0.

	var f4 *int = new(int)
	*f4 = 600

	f1 = 100
	var f5 *int = &f1

	var f6 *int = new(int)
	*f6 = 606

	//var numPtr *int = new(int)
	//numPtr++              // compile error. it is not to allow pointer operating
	//numPtr = 0xc0820062d0 // compile error. it is not to allow assign address

	//////////////////////////////////////////////////
	// print
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println("\n")

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println("\n")

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println("\n")

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println("\n")

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println("\n")

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Printf("[f4] value:(%d), address:(%d)", *f4, f4)
	fmt.Printf("[f5] value:(%d), address:(%d)", *f5, f5)
	fmt.Printf("[f6] value:(%d), address:(%d)", *f6, f6)
	fmt.Println("\n")

	fmt.Println("\n")
}
