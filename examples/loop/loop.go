package main

import (
	"fmt"
)

func main() {

	/*
		loop example
	*/

	// let's start

	//////////////////////////////////////////////////
	// for statement

	var a1 int = 0
	for i := 1; i <= 100; i++ {
		a1 += i
	}

	fmt.Printf("from 1 to 100, a1:(%d) \n", a1)

	//////////////////////////////////////////////////
	// for statement , only conditional expression

	var b1 int = 1
	for b1 < 100 {
		b1 *= 2
	}

	fmt.Printf("b1:(%d) \n", b1)

	//////////////////////////////////////////////////
	// for statement , Infinite loop
	// press 'Ctrl + C' to exit
	/*
	       for {
	           println("Infinite loop")
	   	}
	*/

	fmt.Printf("\n")

	//////////////////////////////////////////////////
	// for range statement
	c1 := []string{"James", "David", "Julia"}

	for index, name := range c1 {
		fmt.Printf("index:(%d), name:(%s) \n", index, name)
	}

	//////////////////////////////////////////////////
	// for statement with break, continue, goto

	var d1 int = 1

	for d1 < 10 {

		if d1 == 1 {
			d1 += d1
			continue
		}

		d1++
		if d1 > 5 {
			break
		}

	}

	if d1 == 6 {
		goto END //goto 사용예
	}

	fmt.Printf("d1:(%d) \n", d1)

END:
	fmt.Printf("Here is END\n")

}
