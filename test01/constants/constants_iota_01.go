package main

import (
	"fmt"
)


const (
	_ =iota //0
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)

)

func main () {
	fmt.Println("brnary\t\tdecimal")
	fmt.Println("%b\t",KB)
	fmt.Println("%d\n",KB)
	fmt.Println("%b\t",MB)
	fmt.Println("%d\n",MB)
	fmt.Println("%b\t",GB)
	fmt.Println("%d\n",GB)
}