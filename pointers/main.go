package main

import "fmt"

// accepts value at a, value  at b
func swapVar(a *int, b *int) {
	temp := *a //
	*a = *b
	*b = temp
}

func main() {
	a := 8
	b := 9

	fmt.Println(a, b)
	swapVar(&a, &b)
	fmt.Println(a, b)
	fmt.Print("cmd ", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98 \n")
}
