package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		fmt.Println(v) // 3, 4, 5, 7, 8, 9
	}
}
