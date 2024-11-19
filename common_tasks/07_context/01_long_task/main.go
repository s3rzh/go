package main

import "time"

// Вводные:
// Функция executeTask может зависнуть.
// В ней не предусмотрен механизм отмены.
// Она не принимает Context или канал с событием отмены как аргумент.

func main() {

	// Задача:
	// Для функции executeTask написать обертку executeTaskWithTimeout.
	// Функция executeTaskWithTimeout принимает аргументом тайм-аут,
	// через который функция executeTask будет отменена.
	// Если executeTask была отменена по тайм-ауту, нужно вернуть ошибку

	executeTask()
}

func executeTask() {
	time.Sleep(10 * time.Second)
}
