package main

import (
	"flag"
	"fmt"
)

func main() {

	/*
		function example
	*/

	// let's start

	//////////////////////////////////////////////////
	// flag statement
	// command line
	//		   flag.String (option name, option default vaule, option usage)
	pName := flag.String("name", "", "this is name")                   // save string option
	pAge := flag.Int("age", 0, "this is  age")                         // save integer option
	pMarriage := flag.Bool("marriage", false, "this is marital state") // save boolean option

	flag.Parse() // parse command line option, according to above definition

	if flag.NFlag() == 0 { // command line option is none
		flag.Usage() // show usage
		return
	}

	fmt.Printf("name : %s\n", *pName)
	fmt.Printf("age : %d\n", *pAge)
	fmt.Printf("marriage : %t\n", *pMarriage)

	fmt.Printf("cmdflag done  \n")
}
