package mian

import (
	"fmt"
)

func main() {
	a := 52
	fmt.Println(a)
	fmt.Println(&a)
	//& gives an address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) //* gives the values stored at an address when you have the address

	*b = 43
	fmt.Println(a)

}
