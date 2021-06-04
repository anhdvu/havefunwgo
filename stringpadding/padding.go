package main

import "fmt"

func main() {
	s := "hahaha"
	// pad := 6

	fmt.Printf("'padding %6s'\n", s)
	fmt.Printf("'%4dkm'\n", 10)
	fmt.Printf("'%5v'\n", "john")
	padded := fmt.Sprintf("%12v", "pad")
	fmt.Println(padded)
}
