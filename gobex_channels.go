package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	buffer_messages := make(chan string, 2)

	buffer_messages <- "buffering"
	buffer_messages <- "a channel"

	fmt.Println(<-buffer_messages)
	fmt.Println(<-buffer_messages)
}
