package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "GoodMorning",
		1: "Namsta",
		2: "test",
		3: "will do",
	}
	fmt.Println(myGreeting)
	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)

	} else {
		fmt.Println("That value Doesn't exist")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}
	fmt.Println(myGreeting)
}
