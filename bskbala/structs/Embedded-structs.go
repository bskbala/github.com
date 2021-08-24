package main

import (
	"fmt"
)
//==> Struct
/*
==> Struct is a sequence of named elements ,called fields ,each of which has name and a type.
within the stuct,non blank filed must be Unique
*/
type person struct { 
		first string
		last  string
}

//==> Embedded Struct 
/*
==>  A field declared with a type but no explicit field name is anonymous field  and also called an embedded field
*/
type scretAgent struct{ 
		person
		ltk bool
}

func main() {
	sa1 := scretAgent{
				person:person{
					first: "Saikumar",
					last:  "Sanjay",
			
			},
			ltk : true,
		}
	p2 := person{
		first: "Balam",
		last:  "Mamidi",
	}
	fmt.Println(sa1)
	fmt.Println(p2)

}