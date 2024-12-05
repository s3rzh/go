package main

import (
	"fmt"
	"iter"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	for k, v := range backwards(s) {
		if v == 3 {
			fmt.Println("loop returned break/false")
			break // break convert to false here -> if !yield(i, s[i]) { и итератор прекращает возвращаться новые значения
			// либо continue что позволит взять следуюищй элемент
		}
		fmt.Println(k, ":", v)
		fmt.Println("loop returned continue/true")
		// continue // true here -> if !yield(i, s[i]) - as result !true = false
	}

	// или тоже самое
	fmt.Println()
	//printEl(backwards(s))
}

func backwards(s []int) func(func(int, int) bool) { // equal: func backwards(s []int) iter.Seq2[int, int] {
	//func backwards(s []int) func(func(int) bool) { // equal: // func backwards(s []int) iter.Seq[int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(i, s[i]) { // если от вызывающей стороны (те циклыа for range пришел false он же break)
				fmt.Println("yield returned false")
				return // прерывание интератора
			}
			fmt.Println("yield returned true")
		}
	}
}

func printEl[V any](s iter.Seq2[int, V]) {
	for k, v := range s {
		if k == 3 {
			continue
		}
		fmt.Println(k, ":", v)
		//fmt.Println("loop returned continue/true")
	}
}

// запуск линтера
// golangci-lint run -v

// output:

// 0 : 1
// loop returned continue/true
// yield returned true
// 1 : 2
// loop returned continue/true
// yield returned true
// loop returned break/false
// yield returned false

