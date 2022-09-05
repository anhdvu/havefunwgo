package main

import (
	"flag"
	"fmt"
	"log"
)

type options struct {
	Env    string
	Port   int
	Logger *log.Logger
}

func main() {
	var opts options

	flag.StringVar(&opts.Env, "env", "development", "Environment (development|staging|production)")
	flag.IntVar(&opts.Port, "port", 1234, "API server port")
	flag.Parse()

	opts.Logger = log.Default()
	fmt.Printf("%+v\n", opts)
}
