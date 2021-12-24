package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice)       // it will print all the slices values
	fmt.Println(mySlice[2:4])  // slicing in slice
	fmt.Println(mySlice[2])    //index accesing :accessing by index
	fmt.Println("myString"[2]) //index access :index accessing

}
