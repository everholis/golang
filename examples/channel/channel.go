package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		channel example
	*/

	// let's start

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel statement (Unbuffered Channel) : waitful

	fmt.Printf("[channel_test_1]\n")

	// create integer channel
	ch001 := make(chan int) // channel name is ch001

	go func() {
		fmt.Printf("[channel_test_1] anonimous function 1 called \n")
		time.Sleep(time.Second * 1)
		ch001 <- 100 //send 100 to channel
	}()

	var nValue int
	nValue = <-ch001 // wait until receiving data(100) from channel
	fmt.Printf("[channel_test_1] nValue:(%d)\n", nValue)

	//
	// warning
	/*
		ch002 := make(chan int)
		ch002 <- 1   // deadlock! because no receiver
		fmt.Println("[channel_test_1] ch002:(%d)\n", <-ch002) //deadlock! goroutine is needed
	*/

	fmt.Printf("channel next step ==========\n")

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel statement (Buffered Channel) : unwaitful

	//create integer channel with one buffer
	ch003 := make(chan int, 1)
	fmt.Printf("[channel_test_3] execute ch003 <- 300 \n")
	ch003 <- 300 // fine, one buffer is available
	//fmt.Printf("[channel_test_3] execute ch003 <- 301 \n")
	//ch003 <- 301 	// deadlock! because out of buffer
	fmt.Printf("[channel_test_3] nValue:(%d)\n", <-ch003) // find

	fmt.Printf("channel next step ==========\n")

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel parameter statement

	ch004 := make(chan string, 1)
	MyFunc4a(ch004) //send to channel
	fmt.Printf("[channel_test_4] strValue:(%s) \n", <-ch004)

	ch004 <- "hi_channel"
	MyFunc4b(ch004) //receive from channel

	ch004 <- "bye_channel"
	//ch004 <- "byebye_channel" //error, because of one buffer channel
	MyFunc4b(ch004) //receive from channel
	//MyFunc4b(ch004) // error, because of one buffer channel

	fmt.Printf("channel next step ==========\n")

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel parameter statement
	// "chan<-" send to channel
	// "<-chan" recv from channel

	ch005 := make(chan string, 1)

	SendToChannel(ch005)
	time.Sleep(time.Second * 1)
	RecvFromChannel(ch005)

	//fmt.Printf("[channel_test_5] strValue:(%s) \n", <-ch005) //error, receiver is already exist

	fmt.Printf("channel next step ==========\n")

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel close statement

	ch006 := make(chan int, 2)

	// send data to channel
	ch006 <- 500
	ch006 <- 501

	// close clannel
	close(ch006)

	// if channel is closed, it is not able to send data to channel
	// but it is able to receive data from channel
	//ch006 <- 503  // error because channel is closed,

	// recv data from channel
	fmt.Printf("[channel_test_6] nValue:(%d) \n", <-ch006)
	fmt.Printf("[channel_test_6] nValue:(%d) \n", <-ch006)

	if _, value := <-ch006; !value {
		fmt.Printf("[channel_test_6] no more data \n")
	}

	fmt.Printf("channel next step ==========\n")

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel close monitoring statement
	ch007 := make(chan int, 2)

	// send data to channel
	ch007 <- 1
	ch007 <- 2

	// close channel
	close(ch007)

	//
	// case 1, how to monitoring channel all closing
	/*
		   for {
		       if i, value := <-ch007; value {
				   fmt.Printf("[channel_test_7] nValue:(%d) \n", value)
		       } else {
		           break
		       }
		   }
	*/

	// case 2, how to monitoring channel all closing
	for i := range ch007 {
		fmt.Printf("[channel_test_7] channel is alive, i:(%d) \n", i)
	}

	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	// channel select statement

	ch008a := make(chan bool)
	ch008b := make(chan bool)

	go MyFunc8a(ch008a)
	go MyFunc8b(ch008b)

EXIT:
	for {
		select { // wait for data receving of channel
		case <-ch008a:
			println("run1 완료")
			fmt.Printf("[channel_test_8] MyFunc1 is done \n")

		case <-ch008b:
			fmt.Printf("[channel_test_8] MyFunc2 is done \n")

			// break lable: out of 'for loop', then goto EXIT,
			// then print "[channel_test_8] done"  under "for loop"
			break EXIT //break lable

			// if default is exist, 'select' don't wait channel, then excute default
			//default
			//	fmt.Printf("[channel_test_8] default \n")

		}
	}
	fmt.Printf("[channel_test_8] done \n")

	fmt.Printf("[channel] done \n")
}

func MyFunc4a(chParam chan string) {
	chParam <- "HelloChannel"
	fmt.Printf("[MyFunc4a] called \n")
}

func MyFunc4b(chParam chan string) {
	value := <-chParam
	fmt.Printf("[MyFunc4b] value:(%s) \n", value)
}

func SendToChannel(chParam chan<- string) {
	chParam <- "helloworld"
	// s := <-chParam // error
	fmt.Printf("[SendChannel] called \n")
}

func RecvFromChannel(chParam <-chan string) {
	value := <-chParam
	fmt.Printf("[RecvFromChannel] value:(%s) \n", value)
}

func MyFunc8a(done chan bool) {
	time.Sleep(time.Second * 1)
	done <- true
}

func MyFunc8b(done chan bool) {
	time.Sleep(time.Second * 2)
	done <- true
}
