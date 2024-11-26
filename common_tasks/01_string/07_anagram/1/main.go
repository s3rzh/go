package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CheckIfStringsAreAnagram("anagram", "nagaram")) // true
}

func CheckIfStringsAreAnagram(source string, target string) bool {
	if len(source) != len(target) {
		return false
	}

	sourceArray := []rune(source)
	sort.Slice(sourceArray, func(i, j int) bool {
		return sourceArray[i] < sourceArray[j]
	})

	targetArray := []rune(target)
	sort.Slice(targetArray, func(i, j int) bool {
		return targetArray[i] < targetArray[j]
	})

	for i := 0; i < len(sourceArray); i++ {
		if sourceArray[i] != targetArray[i] {
			return false
		}
	}

	return true
}

// Решение через сортировку. Вариант заключается в том, чтобы отсортировать обе данные строки и просто-напросто проверить, равны ли они после этого друг другу. Если да, то перед нами анаграмма.

// Анаграма - так называют слово, которое содержит все буквы другого слова в том же количестве, но ином порядке.
