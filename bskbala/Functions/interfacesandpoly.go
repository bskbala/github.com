pacakge main

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

func speak (secretagent)  {
	 fmt.Println("test",s.firstName,s.lastName)
}
func main () {
    sa1 := secretagent{
			person:person{
					"test1",
					"test2",
			},
			ltk:true,
			fmt.Println(sa1)

	}
}