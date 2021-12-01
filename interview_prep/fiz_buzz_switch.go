//FizzBuzz is comman programming test 
//The Objective of the Program is print the number from 1 to 100 
// 1.for the Multiples of 3 print fizz instead of numbers
// 2.for the Multiples of 5 print Buzz instead of numbers
// 3.for the Multiples of 3 and 5 print fizzbuzz instead of numbers


package main

import (
	"fmt"
)

func main  () {
	for index :=1; index <100;index++{
		switch   {
		case index%15 ==0:
			fmt.Println("fizzbuzz")
		case index%3 ==0:
			fmt.Println("buzz")
		case index%5 ==0:
			fmt.Println("fizz")
		default:
			fmt.Println("index")
		
			
		}

	}

}
