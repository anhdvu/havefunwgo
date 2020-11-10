package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, e := range s {
		sum += e
	}
	c <- sum
}

func main() {
	// Example 1
	s := []int{1, 3, 5, 7, 11, -5, -6, -7}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	a, b := <-ch, <-ch
	fmt.Println(a + b)

	// Example 2
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// Example 3

}
