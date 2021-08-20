package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 8}
	x = append(x, 88, 99, 66, 55)
	fmt.Println(x)

}
