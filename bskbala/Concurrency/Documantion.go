/*
They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations.
A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.

what is Goroutines?
Goroutines are multiplexed onto multiple OS threads so if one should block ,such as while waiting for i/o
others continue to run .their design hides many of the complexites of thread creation and managment

==>What is Go Statments?
A go Statments starts the Execution  of a funcation  cal as an independent concurrent thread of control,or
goroutine,within the same address space 

*/


package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
