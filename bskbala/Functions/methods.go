package main

import (
	"fmt"
)

type person struct{
	firstName string
	lastName string
}

type secretagent struct{
	person 
	ltk bool
}

func speak (secretagent){

}
func main () {
	sa1 := secretagent{
		person:person{
			"santra",
			"bond",
		},
		ltk: true,
		sa2 := secretagent{
			person:person{
				"mantra",
				"bond",
			},
			ltk: true,
		fmt.Println("sa1")
		fmt.Println("sa2")
		sa1.speak()
		sa2.speak()

	}
}