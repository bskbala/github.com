package main

import (
	"fmt"
)

func main() {
	greeting := make([]string, 3, 5)
	// 3 is length -number of  the elements referred to by the slice
	//5 is the capacity - number of the elements in the Underling array
	// you could also od it Like this

	greeting[0] = "Good Morning!"
	greeting[1] = "west"
	greeting[2] = "jildil"
	greeting = append(greeting, "test")
	fmt.Println(greeting[3])
}
