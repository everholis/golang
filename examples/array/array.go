package main

import (
	"fmt"
)

func main() {

	/*
		array example
	*/

	// let's start

	//////////////////////////////////////////////////
	// array statement

	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	for index, a1 := range arr1 {
		fmt.Printf("index:(%d), a1:(%d) \n", index, a1)
	}

	fmt.Println(" \n")

	//////////////////////////////////////////////////
	// array statement

	var arr2 = [3]int{11, 22, 33}

	for index, a1 := range arr2 {
		fmt.Printf("index:(%d), a1:(%d) \n", index, a1)
	}

	fmt.Println(" \n")

	//////////////////////////////////////////////////
	// array statement

	var arr3 = [...]int{111, 222, 333}

	for index, a3 := range arr3 {
		fmt.Printf("index:(%d), a3:(%d) \n", index, a3)
	}

	fmt.Println(" \n")

	//////////////////////////////////////////////////
	// array statement

	var mutiarr4 [2][3]int
	mutiarr4[0][0] = 1
	mutiarr4[0][1] = 2
	mutiarr4[0][2] = 3
	mutiarr4[1][0] = 4
	mutiarr4[1][1] = 5
	mutiarr4[1][2] = 6

	for indexi, arr4 := range mutiarr4 {
		for indexj, a4 := range arr4 {

			fmt.Printf("indexi:(%d), indexj:(%d), a4:(%d) \n", indexi, indexj, a4)
		}

	}

	fmt.Println(" \n")

	//////////////////////////////////////////////////
	// array statement

	var mutiarr5 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for indexi, arr5 := range mutiarr5 {
		for indexj, a5 := range arr5 {

			fmt.Printf("indexi:(%d), indexj:(%d), a5:(%d) \n", indexi, indexj, a5)
		}

	}

	fmt.Println(" \n")

}
