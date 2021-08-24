package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Saikumar",
		last:  "Sanjay",
	}
	p2 := person{
		first: "Balam",
		last:  "Mamidi",
	}
	fmt.Println(p1)
	fmt.Println(p2)

}
