package main

import (
	"fmt"
)

type FuncTypeCallback func(int) string

func main() {

	/*
		function example
	*/

	// let's start
	var paramstr string = ""

	var rtstr string = ""
	var rtnum int = 0
	var rtbool bool = false

	_, _, _ = rtnum, rtstr, rtbool

	//////////////////////////////////////////////////
	// function statement
	// special feature: func 'function name' ( param data )  retun type
	// ex) func myfunc( myparam int ) string

	MyFunc1()
	fmt.Printf("MyFunc1 done\n")

	paramstr = "HelloWorld Param"
	MyFunc2(paramstr)
	fmt.Printf("MyFunc2 return, paramstr:(%s)\n", paramstr)

	paramstr = "HelloWorld Param"
	MyFunc3(&paramstr)
	fmt.Printf("MyFunc3 return, paramstr:(%s)\n", paramstr)

	rtstr = MyFunc4(401)
	fmt.Printf("MyFunc4 return, rtstr:(%s) \n", rtstr)

	rtstr = MyFunc5(501, 502, 503)
	fmt.Printf("MyFunc5 return, rtstr:(%s) \n", rtstr)

	rtstr, rtnum, rtbool = MyFunc6()
	fmt.Printf("MyFunc6 return, rtstr:(%s), rtnum:(%d), rtbool:(%t) \n", rtstr, rtnum, rtbool)

	var rtstr7 string = ""
	var rtnum7 int = 0
	var rtbool7 bool = false
	rtstr7, rtnum7, rtbool7 = MyFunc7()
	fmt.Printf("MyFunc7 return, rtstr7:(%s), rtnum7:(%d), rtbool7:(%t) \n", rtstr7, rtnum7, rtbool7)

	//////////////////////////////////////////////////
	// Anonymous Function statement
	MyFunc8 := func(paramnum int) string { //익명함수 정의
		fmt.Printf("MyFunc8 ..... paramnum:(%d) \n", paramnum)
		return "MyFunc8 done"
	}

	rtstr = MyFunc8(801) //call Anonymous Function

	fmt.Printf("MyFunc8 return, rtstr:(%s) \n", rtstr)

	//////////////////////////////////////////////////
	// function prototype  statement

	rtstr = MyFunc9(Callback)
	fmt.Printf("MyFunc9 return, rtstr:(%s) \n", rtstr)

}

//////////////////////////////////////////////////
// function statement - base

func MyFunc1() {

	fmt.Printf("MyFunc1 ..... \n")

}

//////////////////////////////////////////////////
// function statement - Pass By Value
func MyFunc2(paramstr string) {

	fmt.Printf("MyFunc2 ..... paramstr:(%s) \n", paramstr)
}

//////////////////////////////////////////////////
// function statement - Pass By Reference
func MyFunc3(pParamstr *string) {

	fmt.Printf("MyFunc3 ..... paramstr:(%s) \n", *pParamstr)
	*pParamstr = "Welcome"

}

//////////////////////////////////////////////////
// function statement - parameter and return
func MyFunc4(paramnum int) string {

	fmt.Printf("MyFunc4 ....., paramnum:(%d) \n", paramnum)
	return "MyFunc4 done"
}

//////////////////////////////////////////////////
// function statement - variable parameters
func MyFunc5(paramnums ...int) string {

	for index, param := range paramnums {
		fmt.Printf("MyFunc5 ....., index:(%d), param:(%d) \n", index, param)
	}
	return "MyFunc5 done"
}

//////////////////////////////////////////////////
// function statement - multiful return value
func MyFunc6() (string, int, bool) {

	var rtstr string = "HelloWorld"
	var rtnum int = 601
	var rtbool bool = true

	return rtstr, rtnum, rtbool
}

//////////////////////////////////////////////////
// function statement - return variables definition
func MyFunc7() (rtstr7 string, rtnum7 int, rtbool7 bool) {

	rtstr7 = "HelloWorld"
	rtnum7 = 701
	rtbool7 = true

	return
}

//////////////////////////////////////////////////
// function prototype  statement
func Callback(paramnum int) string {

	fmt.Printf("Callback ....., paramnum:(%d) \n", paramnum)
	return "Callback done"
}

func MyFunc9(cb FuncTypeCallback) string {

	var rtstr = cb(901)
	fmt.Printf("MyFunc9 ....., rtstr:(%s) \n", rtstr)

	return "MyFunc9 done"
}
