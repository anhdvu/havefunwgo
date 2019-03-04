package main

import (
	"fmt"
)

// var x int = 42
// var y string = "James Bond"
//  var z bool = true

type dace int

var x dace
var y int

func main() {
	// x := 42
	// y := "James Bond"
	// z := true

	// fmt.Println(x, y, z)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// s := fmt.Sprintf("%v %v %v", x, y, z)
	// fmt.Println(s)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	a := "hahahaha"
	fmt.Println(len(a))
	fmt.Printf("%T\n", len(a))

	y = -1
	fmt.Println(y)

	z1 := []string{}
	fmt.Println(z1)
	fmt.Printf("%T\n", z1)

	z2 := [6]string{"zero", "one", "two"}
	fmt.Println(z2)
	fmt.Printf("%T\n", z2)
	stest := ""

	for _, value := range z2 {
		stest += value
	}
	fmt.Println(stest)

	z3 := map[string]string{"Sunday": "1", "Monday": "2", "Tuesday": "3", "Wednesday": "4", "Thursday": "5", "Friday": "6", "Saturday": "7"}
	fmt.Println(z3)
	fmt.Printf("%T\n", z3)

}

func foo() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
