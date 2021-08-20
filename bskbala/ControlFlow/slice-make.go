package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0]=42
	x[9]=999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,1234)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,12346)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x=append(x,12345)
	
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
