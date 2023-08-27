package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	examineRune := func(s rune) {
		if s == 't' {
			fmt.Println("found t")
		}
	}

	const s = "ଅନୁଚ୍ଛେଦ"
	fmt.Println(s)
	fmt.Println(utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("\n%#U is at index %d \n", runeValue, idx)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])

		fmt.Printf("%#U at index %d \n ", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}
