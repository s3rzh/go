package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "привет"

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}

// Output: 6
