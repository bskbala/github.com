package main

import (
	"fmt"
)

func main() {
	s := greet("sai kumar", "balam")
	fmt.Println(s)

}

func greet(firstname, lastname string) string {
	return fmt.Sprint(firstname, lastname)
}
