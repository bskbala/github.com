package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("sai kumar", "balam"))

}
func greet(firstname, lastname string) (string, string) {
	return fmt.Sprint(firstname, lastname), fmt.Sprint(firstname, lastname)
}
