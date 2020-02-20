package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	a := 0
	b := 1
	return func(i int) int {
		if i == 0 {
			return a
		} else if i == 1 {
			return b
		} else {
			c := a + b
			a = b
			b = c
			return b
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
