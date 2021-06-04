package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, _ := os.Getwd()
	dir := filepath.Dir(cwd)
	upperDir := filepath.Dir(dir)
	list := filepath.SplitList(cwd)
	fmt.Println("1 ", cwd)
	fmt.Println("2 ", dir)
	fmt.Println("3", upperDir)
	fmt.Println(filepath.Join(upperDir, "historical_data.txt"))
	fmt.Println("5 ", filepath.Base(cwd))
	fmt.Println(filepath.Abs(cwd))
	fmt.Println(list[0])
	fmt.Println(filepath.Base(cwd))
}
