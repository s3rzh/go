package main

import (
	"fmt"
)

// что выведет программа?

func main() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)     // 1 2 3
	process(slice)
	fmt.Println(slice)     // 3 нуля
	fmt.Println(slice[:6]) // 6 нулей (последний не включительно)
}

func process(slice []int) {
	for index := range slice {
		slice[index] = 0
	}
}

// Первый вывод: fmt.Println(slice)
// Вывод: [1 2 3]
// Изначально в слайс добавляются элементы 1, 2 и 3.

// Второй вывод: fmt.Println(slice) после process(slice)
// Вывод: [0 0 0]
// Функция process(slice) заменяет каждый элемент слайса на 0, изменяя его на месте.

// Третий вывод: fmt.Println(slice[:6])
// Вывод: [0 0 0 0 0 0]
// Слайс изначально создан с длиной 0 и емкостью 1000, поэтому slice[:6] имеет доступ к первым шести элементам подлежащего массива. Так как process(slice) изменяет только первые три элемента, остальные элементы остаются нулевыми, и выводится [0 0 0 0 0 0].
