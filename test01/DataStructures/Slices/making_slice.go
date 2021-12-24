package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 8, 99, 66}
	for i, value := range mySlice {
		fmt.Println(i, "-", value)

	}
}
