package main
import (

	"fmt"
)
func main() {
	m:= map[string]int{
		"james":32,
		"Miss MoneyPenny": 27,

	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["Barnabas"])

	//map Delete
	delete(m,"james")

}
