package main
import (
	"fmt"
)
func main() {
	p1 :=struct{
		firstName  string
		lastName string
		age int
	}{
		firstName: "saikumar",
		lastName: "Balam",
		age: 24,
	}
	fmt.Println(p1.firstName,p1.lastName,p1.age)	 
}  