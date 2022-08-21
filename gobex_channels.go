package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	//basic
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	//buffering
	buffer_messages := make(chan string, 2)

	buffer_messages <- "buffering"
	buffer_messages <- "a channel"

	fmt.Println(<-buffer_messages)
	fmt.Println(<-buffer_messages)

	//synchronization

	done := make(chan bool, 1)
	go worker(done)
	<-done
}
