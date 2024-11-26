package main

import (
	"fmt"
)

// Данное решение не подайдёт для многобайтных символов (напр. кириллице, иероглифов)

func main() {
	fmt.Println(IsPalindrome("Hello")) // false
	fmt.Println(IsPalindrome("oko"))   // true
}

func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if str[0] != str[len(str)-1] {
		return false
	}

	return IsPalindrome(str[1 : len(str)-1])
}

// Палиндром — слово, предложение или последовательность символов, которая абсолютно одинаково читается как в привычном направлении, так и в обратном. К примеру, “Anna” — это палиндром, а “table” и “John” — нет.
