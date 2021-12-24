package main

import (
	"fmt"
)

func main() {
	fmt.Println(give("test", "me"))
}

func give(firstname string, lastname string) string {
	return fmt.Sprint(firstname, lastname)
}
