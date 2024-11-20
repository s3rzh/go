package main

import (
	"fmt"
)

// что выведет программа и почему?
func main() {
	first := []int{10, 20, 30, 40}
	second := make([]*int, len(first))
	for i, v := range first {
		// v := v // до версии go 1.22
		second[i] = &v
	}
	fmt.Println(*second[0], *second[1]) // 10, 20
}
