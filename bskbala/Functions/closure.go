// we want to enclose the variable  in some code keep the scope more narrow 

package main
 

import (
	"fmt"
)

var x int 

func main () {
	fmt.Println(x)
	x++
	{
		y:=42
		fmt.Println(y)
	}
	fmt.Println(x)
	foo()
	fmt.Println(x)
}
 
func foo () { 
	fmt.Println("helloo")
	x++
}