package main

import ("fmt"
)

func main () {
	myGreeting := map[string]string {
		"zero": "Good Morning",
		"one": "Namsta",
		"two":"sai",
		"three":"red",
	}
	fmt.Println(myGreeting)
	delete(myGreeting,"two")
	fmt.Println(myGreeting)
}