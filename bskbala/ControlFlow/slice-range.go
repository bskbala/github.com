package main
 import (
	  "fmt"
 )
 func main () {
 x := []int{4,5,7,8,42}
 fmt.Println(x[0])
 fmt.Println(x[1])
 fmt.Println(x[2])
 fmt.Println(x[3])
 fmt.Println(x[4])

  for i,v := range x{
	  fmt.Println(i,v)
  }
 }