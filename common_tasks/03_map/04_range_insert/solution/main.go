package main

import (
	"fmt"
)

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true}

	m2 := copyMap(m)

	for k, v := range m {
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	copiedMap := make(map[int]bool, len(m))
	for k, v := range m {
		copiedMap[k] = v
	}
	return copiedMap
}

// решение через копию мапы, сначало скопировать (доп. функция), потом читать из исходной, а писать в скопированную.
