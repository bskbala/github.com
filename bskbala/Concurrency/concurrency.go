/*What is ConCurrency ?
 Concurrency is an ability of a program to do multiple things at the same time.
This means a program that have two or more tasks that run individually of each other, at about the same time, but remain part of the same program.
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	foo()
	bar()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
