package main

import (
	"log"
	"time"

	"github.com/anhdvu/havefunwgolang/list-directory/store"
)

func main() {
	p := "../data"

	s, err := store.NewTextFileStore(p)
	if err != nil {
		log.Fatal(err)
	}
	s.ListFiles()
	time.Sleep(30 * time.Second)
}
