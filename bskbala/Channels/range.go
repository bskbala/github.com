package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//send
	go foo(c)

	//receive
	for v :=range c  {
		fmt.Println(v)
	}


	fmt.Println("About To Exit")
}

//send
func foo(c chan<- int) {
	for i :=0; i <100 ;i++{
		c <- 42
	}
	close (c)
}



