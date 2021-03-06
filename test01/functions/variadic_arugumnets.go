package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, V := range sf {
		total += V
	}
	return total / float64(len(sf))
}
