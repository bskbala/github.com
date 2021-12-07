package main

//keyword called defer

import "fmt"

func hello() {
	fmt.Println("hello")

}
func world() {
	fmt.Println("world")
}

func main() {
	world()
	hello()
}
