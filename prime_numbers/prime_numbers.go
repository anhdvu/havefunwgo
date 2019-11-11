package main

import (
	"fmt"
	"math"
)

// Beginning of main function to execute this script
func main() {
	fmt.Println(primeNumbers(10))
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number >= 2 {
		for f := 2; f < int(math.Sqrt(float64(number)))+1; f++ {
			if number%f == 0 {
				return false
			}
		}
	}
	return true
}

func primeNumbers(n int) []int {
	result := make([]int, 0)

	for i := 0; i <= n; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}

	return result
}
