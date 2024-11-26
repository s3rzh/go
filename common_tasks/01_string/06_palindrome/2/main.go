package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("Hello")) // false
	fmt.Println(IsPalindrome("oko"))   // true
}

func IsPalindrome(str string) bool {
	reversedStr := strings.Builder{}

	for i := len(str) - 1; i >= 0; i-- {
		reversedStr.WriteByte(str[i])
	}

	return str == reversedStr.String()
}

// Палиндром — слово, предложение или последовательность символов, которая абсолютно одинаково читается как в привычном направлении, так и в обратном. К примеру, “Anna” — это палиндром, а “table” и “John” — нет.
