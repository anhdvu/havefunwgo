package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeNumbers(53))
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for f := 2; f < int(math.Sqrt(float64(number)))+1; f++ {
		if number%f == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(n int) []int {
	result := make([]int, 1)

	for i := 0; i <= n; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}

	return result
}
