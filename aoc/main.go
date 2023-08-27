package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	n, err := os.ReadFile("day1.txt")
	check(err)
	fmt.Println(string(n))
}
