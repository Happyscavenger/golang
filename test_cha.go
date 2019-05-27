package main

import (
	"fmt"
	"time"
)

func server11(ch chan string) {
	ch <- "I'm from server1"
}

func server22(ch chan string) {
	ch <- "I'm from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server11(output1)
	go server22(output2)
	time.Sleep(time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	default:
		fmt.Println("I'm default")
	}

}
