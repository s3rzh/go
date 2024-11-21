package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// add timeout to avoid long waiting
func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)
		
	chanForResp := make(chan int)
	go RPCCall(chanForResp)

	select {
	case <-ctx.Done(): // спустя 2е секунды в канал, который возврщает метод Done будет отправлено значение (пустая структура struct{}{})
		fmt.Println("timeout")
	case result := <-chanForResp:
		fmt.Println(result)
	}
}

func RPCCall(ch chan<- int) {
	// sleep 0-4 sec
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))

	ch <- rand.Int()
}

// Задачка про таймауты
// Мы создаем канал chanForResp в который мы ждем ответ из функции RPCCall
// RPCCall имитирует http вызов в какой-нибудь сервис
// в нем мы создаем задержку 0-4 sec
// мы ожидаем вызов функции в 2 секунды, больше нас не устраивает
// мы планируем передавать контекст и отменять его
// отмена контекста мы можем обрабатывать по разному
// допустим мы начали делать вызов, но что-то изменилось
// в этом случае мы прекращаем работу (запрос к сервису или БД)

// у контекста есть метод ctx.Done(), он возвращает канал из которого
// мы что-то прочтем если его отменили, если его не отменили, то ничего мы прочитать не сможем

// создадим контекст с таймаутом 2 секунды,
// наша функция RPCCall создает задержку от 0 до 4 секунд
// и в зависимости от этого мы ожидаем разный результат
// создадим структуру resp для вывода результата

// в этой задаче у нас 2 решения
// 1. select в функции main (например, когда у нас нет возможности изменить вызываемую функцию)
// 2. select в функции RPCCall
