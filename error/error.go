package main

import (
	"fmt"
	"math"
)

// Sqrt function that returns the square root or an error
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Return the custom error for negative input
	}
	return math.Sqrt(x), nil // Return the square root and no error
}

// Define a custom error type for negative square roots
type ErrNegativeSqrt float64

// Implement the Error method for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
