package main

import (
	"fmt"
)

func main() {
	primeNumbers(50)
}

func primeNumbers(n int) {
	var check bool
	for i := 2; i < n; i++ {
		check = true
		if i > 2 {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					check = false
					break
				}
			}
		}
		if check {
			fmt.Println(i)
		}
	}
}
