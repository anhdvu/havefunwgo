package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer1 expired!")

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	if stop2 := timer2.Stop(); stop2 {
		fmt.Println("timer2 has been stopped!")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("main thread ended.")
}
