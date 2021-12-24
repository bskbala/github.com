package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wensday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
}
