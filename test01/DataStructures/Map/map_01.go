package main

import (
	"fmt"
)

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["tim"] = "Good Morning."
	myGreeting["Jenny"] = "Namsta."

	fmt.Println(myGreeting)
}
