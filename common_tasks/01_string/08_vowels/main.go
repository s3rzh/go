package main

import (
	"fmt"
)

func main() {
	fmt.Println(countVowels("Hello")) // 2
}

func countVowels(s string) int {
	count := 0
	for _, char := range s { // в char у нас runa возвращаются, а индексы игнорим.
		switch char {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			count++
		}
	}
	return count
}

// Функция, подсчитывающая количество гласных букв в строке.
