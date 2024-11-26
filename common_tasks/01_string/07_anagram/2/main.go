package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CheckIfStringsAreAnagram("anagram", "nagaram")) // true
}

func CheckIfStringsAreAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sourceArray := []byte(s)
	sort.Slice(sourceArray, func(i, j int) bool {
		return sourceArray[i] < sourceArray[j]
	})

	targetArray := []byte(t)
	sort.Slice(targetArray, func(i, j int) bool {
		return targetArray[i] < targetArray[j]
	})

	return bytes.Equal(sourceArray, targetArray)
}

// с конвертацией в байты
