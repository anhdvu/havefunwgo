package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("Tick at %+v\n", t)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- struct{}{}
	fmt.Println("ticker stopped.")
}
