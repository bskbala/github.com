package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//send
	go foo(c)

	//receive

	go bar(c)

	fmt.Println("About To Exit")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
