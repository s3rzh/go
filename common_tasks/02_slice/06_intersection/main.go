package main

import (
	"fmt"
)

// На вход подаются два неупорядоченных массива любой длины.
// Необходимо написать функцию, которая возвращает пересечение массивов
func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		counter[elem]++
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}

func main() {

	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}

// Можно решить сортировкой, за более долгое время, но без выделения дополнительной памяти. 
// А можно выделить дополнительную память и решить за линейное время.

// Надо посчитать количество появлений элементов первого массива (лучше брать тот, что покороче) — используем для этого словарь. 
// Потом пройтись по второму массиву и вычитать из словаря те элементы, которые есть в нем. 
// По ходу добавляем в результат те элементы, у которых частота появлений больше нуля.
