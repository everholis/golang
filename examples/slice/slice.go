package main

import (
	"fmt"
)

func main() {

	/*
		slice example
	*/

	// let's start

	//////////////////////////////////////////////////
	// slice statement

	var slice1 []int
	slice1 = []int{1, 2, 3}

	for index, a1 := range slice1 {
		fmt.Printf("index:(%d), a1:(%d) \n", index, a1)
	}

	fmt.Printf("slice1 len:(%d), cap:(%d) \n\n", len(slice1), cap(slice1))

	//////////////////////////////////////////////////
	// slice statement

	var slice2 = []int{1, 2, 3}
	slice2[0] = 201
	slice2[1] = 202
	slice2[2] = 203

	for index, a2 := range slice2 {
		fmt.Printf("index:(%d), a2:(%d) \n", index, a2)
	}

	fmt.Printf("slice2 len:(%d), cap:(%d) \n\n", len(slice2), cap(slice2))

	//////////////////////////////////////////////////
	// slice statement

	var slice3 []int

	if slice3 == nil {
		fmt.Println("slice3 is nil")
	}

	fmt.Println("\n")

	//////////////////////////////////////////////////
	// slice statement - use make inner function

	slice3 = make([]int, 3, 5) // Length: 3, Capacity: 5
	slice3[0] = 301
	slice3[1] = 302
	slice3[2] = 303
	//slice3[3] = 304	// error, out of range

	for index, a3 := range slice3 {
		fmt.Printf("index:(%d), a3:(%d) \n", index, a3)
	}

	fmt.Printf("slice3 len:(%d), cap:(%d) \n\n", len(slice3), cap(slice3))

	//////////////////////////////////////////////////
	// slice statement - Sub slice
	var slice4 []int
	slice4 = []int{0, 1, 2, 3, 4, 5}
	subslice41 := slice4[2:5] // index 2 <= x < index 5 , len:3 , cap:4
	subslice42 := slice4[:2]  // 	  		x < index 2 , len:2 , cap:6
	subslice43 := slice4[2:]  // index 2 <= x			, len:4 , cap:4

	for index, a4 := range slice4 {
		fmt.Printf("index:(%d), a4:(%d) \n", index, a4)
	}
	fmt.Printf("slice4 len:(%d), cap:(%d) \n\n", len(slice4), cap(slice4)) // len: 6, cap:6

	for index, a41 := range subslice41 {
		fmt.Printf("index:(%d), a41:(%d) \n", index, a41)
	}
	fmt.Printf("subslice41 len:(%d), cap:(%d) \n\n", len(subslice41), cap(subslice41)) // len: 4, cap:6

	for index, a42 := range subslice42 {
		fmt.Printf("index:(%d), a42:(%d) \n", index, a42)
	}
	fmt.Printf("subslice42 len:(%d), cap:(%d) \n\n", len(subslice42), cap(subslice42)) // len: 2, cap:6

	for index, a43 := range subslice43 {
		fmt.Printf("index:(%d), a43:(%d) \n", index, a43)
	}
	fmt.Printf("subslice43 len:(%d), cap:(%d) \n\n", len(subslice43), cap(subslice43)) // len: 4, cap:4

	//////////////////////////////////////////////////
	// slice statement - append
	var slice5 []int
	slice5 = []int{0, 1}
	slice51 := append(slice5, 2)        // 0, 1, 2
	slice52 := append(slice51, 3, 4, 5) // 0, 1, 2, 3, 4, 5

	for index, a51 := range slice51 {
		fmt.Printf("index:(%d), a51:(%d) \n", index, a51)
	}
	fmt.Printf("slice51 len:(%d), cap:(%d) \n\n", len(slice51), cap(slice51)) // len: 3, cap:4

	for index, a52 := range slice52 {
		fmt.Printf("index:(%d), a52:(%d) \n", index, a52)
	}
	fmt.Printf("slice52 len:(%d), cap:(%d) \n\n", len(slice52), cap(slice52)) // len: 5, cap:6

	//////////////////////////////////////////////////
	// slice statement - capacity
	// if over capacity, Capacity is doubled than before

	var slice6 []int
	slice6 = make([]int, 0, 3) // len: 5, cap:6

	for i := 1; i <= 15; i++ {
		slice6 = append(slice6, i)
		fmt.Printf("slice6 index:(%d), len:(%d), cap:(%d) \n", i, len(slice6), cap(slice6))
	}

	fmt.Println(slice6)
	fmt.Println("\n")

	//////////////////////////////////////////////////
	// slice statement - merge
	// notice: when collection is merge, must use ellipsis(...)

	slice71 := []int{0, 1, 2}
	slice72 := []int{3, 4, 5}
	var slice7 []int
	slice7 = append(slice71, slice72...) //must be ...

	for index, a7 := range slice7 {
		fmt.Printf("index:(%d), a7:(%d) \n", index, a7)
	}
	fmt.Println(" \n")

	//////////////////////////////////////////////////
	// slice statement - copy
	// copy(target, source)

	var slice81 = []int{0, 1, 2}
	var slice82 = []int{3, 4, 5}
	var slice83 = make([]int, len(slice82), cap(slice82))

	copy(slice82, slice81)
	copy(slice83, slice81)

	for index, a82 := range slice82 {
		fmt.Printf("index:(%d), a82:(%d) \n", index, a82)
	}

	for index, a83 := range slice83 {
		fmt.Printf("index:(%d), a83:(%d) \n", index, a83)
	}

}
