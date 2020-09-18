package main

import (
	"fmt"
	"math/rand"
)

func UNUSED(x ...interface{}) {}

func main() {

	/*
		conditional example
	*/

	// let's start

	//////////////////////////////////////////////////
	// if statement

	var a1 int = 0

	if a1 == 1 { // <==  if is need '{' at same line
		fmt.Println("a1 is one")
	}

	if a1 == 2 {
		fmt.Println("a1 is two")
	} else if a1 == 3 {
		fmt.Println("a1 is three")
	} else {
		fmt.Println("a1 is other")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// switch statement
	// notice : switch of golang is not break, switch is done after case

	var b1 = 1

	switch b1 {
	case 1:
		fmt.Println("b1 is one")
	case 2:
		fmt.Println("b1 is two")
	case 3, 4:
		fmt.Println("b1 is three or four")
	default:
		fmt.Println("b1 is default")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// switch statement

	var c1 = 1

	switch c1 {
	case 1:
		fmt.Println("c1 is one")
	case 2:
		fmt.Println("c1 is two")
	case 3, 4:
		fmt.Println("c1 is three or four")
	default:
		fmt.Println("c1 is default")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// Expressionless switch

	var d1 = 1

	switch { // expressionless is omitted
	case d1 <= 100 && d1 >= 90:
		println("d1 is A")
	case d1 >= 80:
		println("d1 is B")
	case d1 >= 70:
		println("d1 is C")
	case d1 >= 60:
		println("d1 is D")
	default:
		println("d1 is F")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// Expressioned switch

	switch d2 := rand.Intn(100); { // expression
	case d2 <= 100 && d2 >= 90:
		println("d2 is A")
	case d2 >= 80:
		println("d2 is B")
	case d2 >= 70:
		println("d2 is C")
	case d2 >= 60:
		println("d2 is D")
	default:
		println("d2 is F")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// switch statement

	var e1 = 1
	var i interface{} = e1

	switch i.(type) {
	case int:
		println("type of e1 is int")
	case bool:
		println("type of e1 is bool")
	case string:
		println("type of e1 is string")
	default:
		println("type of e1 is unknown")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// switch statement
	// fallthrough make switch statement same to c++
	var f1 = 1

	switch f1 {
	case 1:
		fmt.Println("f1 is less than one")
		fallthrough
	case 2:
		fmt.Println("f1 is less than two")
		fallthrough
	case 3:
		fmt.Println("f1 is less than three")
		fallthrough
	default:
		fmt.Println("f1 is less than default")
	}

	fmt.Println("\n")

}
