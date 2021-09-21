/*What is ConCurrency ?
 Concurrency is an ability of a program to do multiple things at the same time.
This means a program that have two or more tasks that run individually of each other, at about the same time, but remain part of the same program.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // wg is package scope

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

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
