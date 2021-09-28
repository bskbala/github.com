package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("HELLO")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
