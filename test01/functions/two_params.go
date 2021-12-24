package main

import (
	"fmt"
)

func main() {
	greet("sai kumar", "Balam")       //Augrments
	greet("satayanarayana ", "Balam") //Augrments
	greet("shyamala", "Balam")        //Augrments

}

func greet(firstname, lastname string) { //parameters
	fmt.Println(firstname, lastname)
}
