package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func fastServer1(ch chan string) {
	ch <- "from fast server1"
}

func fastServer2(ch chan string) {
	ch <- "from fast server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	/* output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	} */

	// ch1 is always empty, runtime error will be thrown if no default case is set
	ch1 := make(chan string)
	select {
	case <-ch1:
	default:
		fmt.Println("default case executed")
	}

	output1 := make(chan string)
	output2 := make(chan string)
	go fastServer1(output1)
	go fastServer2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// an empty select will block forever
	// select {}
}
