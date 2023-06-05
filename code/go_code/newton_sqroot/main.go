package main

import (
	"fmt"
	"math"
)

func Sqrt(number_to_square float64, max_iterations int) float64 {
	// newtonian way to calculate sqrt.
	guess := float64(1) // seeding with the value 1
	for i := 1; i < max_iterations; i++ {
		// divide number_to_square by guess, add guess, divide by 2, use outcome as new guess and try again
		guess = 0.5 * (guess + (number_to_square / guess))
		fmt.Println(guess)
	}
	return guess
}

func main() {
	x := float64(19)
	lim := 19
	outcome_n := Sqrt(x, lim)
	outcome_m := float64(math.Sqrt(x))
	fmt.Println("-----------------------------------")
	fmt.Printf("Newton says:      %v \n", outcome_n)
	fmt.Println("-----------------------------------")
	fmt.Printf("math.Sqrt() says: %v \n", outcome_m)
}
