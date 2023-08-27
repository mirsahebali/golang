package main

import "fmt"

func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	s := seq()
	fmt.Println(s())

	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}
