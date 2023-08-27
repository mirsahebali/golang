package main

import "fmt"

func main() {
	a := vF("1", "3", "4", "5")
	fmt.Println(a)
}

func vF(nums ...string) []string {
	return nums
}
