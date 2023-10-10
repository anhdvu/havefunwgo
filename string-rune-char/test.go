package main

import (
	"fmt"
	"strings"
)

func main() {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	i := strings.IndexRune(alphabet, 'é')
	fmt.Println(i)
	fmt.Println(string(alphabet[i+5]))

	r := []rune(alphabet)
	r[0] = 'á'

	fmt.Println(r)

	fmt.Println(string(r))
}

// func getOffsetChar(c rune, offset int) string {
// 	const alphabet = "abcdefghijklmnopqrstuvwxyz"

// }
