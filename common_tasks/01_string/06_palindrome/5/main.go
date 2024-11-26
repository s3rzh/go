package main

import (
	"fmt"
	"unicode/utf8"
)

// Тут есть panic при кориллице напрмиер!!

func main() {
	fmt.Println(IsPalindrome("Hello")) // false
	fmt.Println(IsPalindrome("oko"))   // true
}

func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	firstRune, _ := utf8.DecodeRuneInString(str)
	lastRune, _ := utf8.DecodeLastRuneInString(str)

	if firstRune != lastRune {
		return false
	}

	return IsPalindrome(str[utf8.RuneLen(firstRune) : utf8.RuneCountInString(str)-utf8.RuneLen(lastRune)])
}


// Палиндром — слово, предложение или последовательность символов, которая абсолютно одинаково читается как в привычном направлении, так и в обратном. К примеру, “Anna” — это палиндром, а “table” и “John” — нет.
