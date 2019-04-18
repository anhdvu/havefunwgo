package main

import (
	"fmt"
)

func main() {
	fmt.Println(primeNumbers(53))
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for f := 2; f < number; f++ {
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
