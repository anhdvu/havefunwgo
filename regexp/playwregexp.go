package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type marker struct {
	time     string
	position string
	color    string
	shape    string
	text     string
}

func main() {
	re := regexp.MustCompile(`(?m)^\d+.*`)
	f := "records.txt"
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}
	matches := re.FindAll(raw, -1)

	l := len(matches)
	last := string(matches[l-1])
	parts := strings.Split(last, `: `)
	tmpTime := parts[0][:16]
	lastDigit, _ := strconv.ParseInt(string(tmpTime[len(tmpTime)-1]), 0, 0)
	if lastDigit < 5 {

	}

	parts2 := strings.Split(parts[1], " ")
	for _, part := range parts2 {
		fmt.Println(string(part))
	}

	// fmt.Printf("%T\n", matches)
	// for _, match := range matches {
	// 	fmt.Println(string(match))
	// }
}
