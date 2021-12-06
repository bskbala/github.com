package main

import (
	"fmt"
)

func vist(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}

}

func main() {
	vist([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

//call back : passing a func as an argument
