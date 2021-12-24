package main 

import ("fmt")

func main () {
	mySlice := []string {"Monday","Tuesday"}
	myOtherSlice :=[]string {"Wensday","Thursday","Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice =append(mySlice[:2],myOtherSlice[3:])
}