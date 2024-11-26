package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckIfStringsAreAnagram("anagram", "nagaram")) // true
}

func CheckIfStringsAreAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sourceMap := make(map[rune]int)
	targetMap := make(map[rune]int)

	for _, letter := range s {
		sourceMap[letter]++
	}

	for _, letter := range t {
		targetMap[letter]++
	}

	for letter, sourceCount := range sourceMap {
		if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
			return false
		}
	}

	return true
}

// Решение через хеш-таблицу.
