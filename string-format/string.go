package main

import "fmt"

type response struct {
	resultCode int
	validCodes []int
	boiler     string
}

func main() {

	r := response{1, []int{1, -9}, `resultCode = %d`}

	o := fmt.Sprintf(r.boiler, r.resultCode)

	fmt.Println(o)
}
