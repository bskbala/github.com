package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 45, 99, 18, 16, 56, 12}
	xs := []string{"james", "Q", "m", "MoneyPenny"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xs)
}
