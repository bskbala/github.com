package main

import (
	"fmt"
)

func main() {
	var x [256]string
	fmt.Println(len(x))

	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}

	for _, v := range x {
		fmt.Println("%v -%T - %v\n", v, v, []byte(v))
		// if i > 50 {
		// 	break
		// }
	}
}
