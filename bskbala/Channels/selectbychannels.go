package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//recevie

	recevie(eve, odd, quit)

}

func recevie(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the Eve Channel:", v)
		case v := <-o:
			fmt.Println("from the odd Channel:", v)
		case v := <-q:
			fmt.Println("from the quit Channel:", v)

			return
		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0

}
