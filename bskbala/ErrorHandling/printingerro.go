/*
You have a few options to choose from when it comes to printing out, or logging, an error message:
fmt.Println()
log.Println()
log.Fatalln()
os.Exit()
log.Panicln()
deferred functions run
can use “recover”
panic(
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no file.txt")
	if err != nil {
		fmt.Println("err happeend", err)
	}
}
