package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(80 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("Concurrency!!!")
		}
	}()
	go say("world")
	say("hello")
}
