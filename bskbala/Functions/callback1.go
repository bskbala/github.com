package main

import (
	"fmt"
)

func main () {
	fmt.Println("Hello Papa")

}

func  sum (xi ... int) int {
	total := 0
	for _, v := range xi {
		total += v

	}
	return total 
}