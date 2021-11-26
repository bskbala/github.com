package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a) // a is printed

	fmt.Println(&a) // address of a is printed

	var b *int = &a //declaring b and getting from a
	fmt.Println(b)  //we will get  the a address here
	fmt.Println(*b) // we will get the a value here this is called  derefrecenig

}
