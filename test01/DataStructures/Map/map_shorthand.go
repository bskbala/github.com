package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Namasta",
	}
	fmt.Println(myGreeting["jenny"])
}
