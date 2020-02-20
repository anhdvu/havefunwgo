package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n < 0 {
		fmt.Println("Wrong input.")
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	prev := 0
	current := 1
	for i := 2; i < n; i++ {
		//fmt.Println("At the begining of loop, prev is ", prev)
		//fmt.Println("At the begining of loop, current is ", current)
		next := current + prev
		//fmt.Println("After the addition, next is ", next)
		prev = current
		//fmt.Println("After the addition, prev is ", prev)
		current = next
		//fmt.Println("After the addition, current is ", current)
	}
	return current + prev
}

func main() {
	for k := 0; k < 9; k++ {
		fmt.Print(fibonacci(k), ", ")
	}
	fmt.Println()
	// 0, 1, 1, 2, 3, 5, 8, 13, 21
}
