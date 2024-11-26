package main

import (
	"bytes"
	"fmt"
)

// Данное решение не подадёт для многобайтных символов

func main() {
	fmt.Println(IsPalindrome("Hello")) // false
	fmt.Println(IsPalindrome("oko"))   // true
}

func IsPalindrome(str string) bool {
	reversedBytes := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		reversedBytes[i] = str[len(str)-i-1]
	}

	return bytes.Equal([]byte(str), reversedBytes)
}

// Палиндром — слово, предложение или последовательность символов, которая абсолютно одинаково читается как в привычном направлении, так и в обратном. К примеру, “Anna” — это палиндром, а “table” и “John” — нет.
