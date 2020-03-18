package main

import (
	"fmt"
)

func main() {
	min := 0
	max := 11
	var numbers []int

	for a := min; a < max; a++ {
		numbers = append(numbers, a)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v is even", number)
		} else {
			fmt.Printf("%v is odd", number)
		}
	}
}
