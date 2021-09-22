/*
==>What is Channels?
=>In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine.
By default channel is bidirectional, means the goroutines can send or receive data through the same channel



package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
