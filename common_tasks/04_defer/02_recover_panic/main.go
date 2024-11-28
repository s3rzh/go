package main

import "fmt"

func main() {
	fmt.Println(division(10, 2))
	fmt.Println(division(10, 0))
	fmt.Println("Выполнение программы продолжается!")
}
func division(x, y int) (n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			n = 0 // Возвращаем из функции division() ноль
		}
	}()
	fmt.Println("Инструкция до деления")
	n = x / y
	fmt.Println("Инструкция после деления")
	return
}
