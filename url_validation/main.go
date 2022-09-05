package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://ttk.dace.dev/wallet"

	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)
}
