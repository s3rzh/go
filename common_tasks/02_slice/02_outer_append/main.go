package main

import (
	"fmt"
)

// что выведет программа?

func main() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)     // 1 2 3
	process(slice)   // передаём в функцию по значению копию структуры слайса, когда слайс изменяется, подлежащий массив изменяется а len и cap снаружи не меняется для вызывающей стороны и на по-прежнему имеет доступ только к 3ем элементам.
	fmt.Println(slice)     // 1 2 3
	fmt.Println(slice[:6]) // 1 2 3 4 5 6
}

func process(slice []int) {
	slice = append(slice, 4, 5, 6)
}


// [1 2 3]
// [1 2 3]
// [1 2 3 4 5 6]


// На 14-ой строке 1 2 3, потому что в функции мы передаем слайс по знамению.
// len у нас останется равным трем. 
// При этом базовый слайс у нас изменился и значения записались.


// На 15-ой строке мы принудительно говорим, чтобы распечать 6 значений базового слайса.
