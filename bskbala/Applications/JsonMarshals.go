package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "bandi",
		last:  "Sai",
		age:   23,
	}
	p2 := person{
		first: "miss",
		last:  "Piss",
		age:   25,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	bs,  err := json.Marshal(people)

	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
