package main

import (
	"fmt"
	"sync"
)

type dep struct {
	c string
}

func (d *dep) print(n int, wg *sync.WaitGroup) {
	// wg.Add(1)
	fmt.Println(fmt.Sprintf("Print %v-%d", d.c, n))
	wg.Done()
}

func main() {
	d := &dep{c: "dace"}
	ints := []int{1, 2, 3, 4, 5, 6, 7}

	var wg sync.WaitGroup
	wg.Add(len(ints))
	for _, n := range ints {
		// wg.Add(1)
		i := n
		go d.print(i, &wg)
	}
	wg.Wait()

	fmt.Println("DONE")
}
