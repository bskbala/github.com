package main

import (
	"fmt"
)

func main() {
	// var rows = 5
	for i := 'A'; i <= 'F'; i++ {
		for j := 'A'; j <= i; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}

}
