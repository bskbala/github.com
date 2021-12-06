package mian

import (
	"fmt"
)

func main() {

	greeting := func() {
		fmt.Println("hellworld")
	}
	greeting()
	fmt.Printf("%T\n", greeting)

}
