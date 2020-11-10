package main

import (
	"fmt"
	"strconv"
)

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(MyInt2Int(i))
}

func MyInt2Int(mi MyInt) int {
	return int(mi)
}

func main() {
	var bla fmt.Stringer = MyInt(123)
	fmt.Println(bla)
}
