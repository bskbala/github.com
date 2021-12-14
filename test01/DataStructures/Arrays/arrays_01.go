package main

import (
	"fmt"
)

func main() {
	var x [256]int
	fmt.Println(len(x))

	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = 1
	}

	for i, v := range x {
		fmt.Println("%v -%T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
