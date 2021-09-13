// When to use the [Pointers]
package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x Before", &x)
	fmt.Println("x Before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}
func foo(y *int) {
	fmt.Println("y Before", *y)
	fmt.Println("y Before", y)
	*y = 43
	fmt.Println("y after", *y)
	fmt.Println("y after", y)
}
