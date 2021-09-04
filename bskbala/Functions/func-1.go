package main

import (
	"fmt"
)

//Syntax func (r Receiver) identifier(Parameters) (returns(s) {....})
func main () {
	fmt.Printf("test")
}

func foo() {
	fmt.Println("Hello form foo")
}

//Everything in go  is pass  by value 
func bar(s string){
	fmt.Println("hello,",s)
}

func woo(st string) {
	return fmt.Sprint("Hello from woo, ",st)

}

