package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	d1 := make(chan string, 1)
	go func() {
		time.Sleep(7 * time.Second)
		d1 <- "result1"
	}()
	select {
	case res := <-d1:
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout 1")
	}

	d2 := make(chan string, 1)
	go func() {
		time.Sleep(9 * time.Second)
		d2 <- "result2"
	}()
	select {
	case res := <-d2:
		fmt.Println(res)
	case <-time.After(11 * time.Second):
		fmt.Println("timeout 2")
	}
}
