package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Bandi")
	fmt.Fprintln(os.Stdout, "Hello hindi")
	io.WriteString(os.Stdout, "Hello wigigng")
}
