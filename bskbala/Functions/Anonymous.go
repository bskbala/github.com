package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println("Anonymous Func Ran")
	}()
	func(x int) {
		fmt.Println("The Meaning of Life :", x)
	}(42)

}

func foo() {
	fmt.Println("foo ran")
}
