package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Namsta",
	}
	myGreeting["King Of test"] = "Rowdy"
	myGreeting["sai "] = "test"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting)
}
