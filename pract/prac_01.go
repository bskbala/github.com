package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("interger Reverse Sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	fmt.Println()

	fmt.Println("String Reverse Sort")
	text := []string{"japan", "uk", "germany"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)

}

/*This the reverse sort Program  to print interger and string

package main

import ("fmt"
		"sort"
)

func main () {
	fmt.Println("Interger Reverse Sort")
	num := [] int {50,20,56,36}
	sort.Sort(sort.Reverse(sort.StringSlice(num)))
	fmt.Println(num)

	fmt.Println()

	fmt.Println("String Reverse Sort")
	text := [] string {"japan","india","uk","Austrila"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)


}

*/
