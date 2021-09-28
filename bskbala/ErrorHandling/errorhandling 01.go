/*What is Error Handling?
==> Errors indicate an unwanted condition occurring in your application.
Let's say you want to create a temporary directory where you can store some files for your application, but the directory's creation fails.
This is an unwanted condition and is therefore represented using an error.
Golang represents errors using the built-in error type,

*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}
