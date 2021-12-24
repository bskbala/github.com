package main

import (
	"fmt"
)

func main() {
	greeting := []string{
		"Good Morning",
		"Please Hit mee hard",
		"test the Shiva",
		"will meeett me in the king",
		"most wanted hit",
	}
	fmt.Print("[1:2]")
	fmt.Println(greeting[2:4])
	fmt.Println(greeting[:2])
	fmt.Print("[:]")
}
