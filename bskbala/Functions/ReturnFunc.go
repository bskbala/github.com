package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)
	x :=func () int {
		return 451
	} ()

	fmt.Println(x)

}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
