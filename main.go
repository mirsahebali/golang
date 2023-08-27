package main

import "fmt"

type Color string

func main() {
	const (
		Red Color = ""
		Green
		Blue
	)
	fmt.Println(Blue)
}
