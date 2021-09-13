package main

import (
	"fmt"
)

func main () {
	fmt.Println(4*3*2*1)
	n := factorial(4)
	fmt.Println(n)

}
func factorial (n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial (4-1) * factorial (3-1) * factorial (2-1)
}
func looFact(n int) int {
	total := 1
	for ; n > 0 ;n -- {
		total *= n
	}
	return total
}