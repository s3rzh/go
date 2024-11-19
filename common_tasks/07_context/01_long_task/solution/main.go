package main

import (
	"context"
	"log"
	"time"
)

func main() {
	timeout := 5 * time.Second
	err := executeTaskWithTimeout(context.Background(), timeout)
	if err != nil {
		log.Fatal(err)
	}
}

func executeTaskWithTimeout(ctx context.Context, timeout time.Duration) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// создаем буферизированный канал
	// если канал не буферизированный и функция отменена по тайм-ауту,
	// то done <- struct{}{} в горутине ниже заблокируется навсегда и горутина останется зависшей
	done := make(chan struct{}, 1)

	go func() {
		executeTask()

		// сигнализируем об окончании executeTask
		done <- struct{}{}

		// закрываем канал
		close(done)
	}()

	// ожидаем выполнение executeTask (слушая канал done) или отмены контекста по тайм-ауту
	select {
	case <-done: // executeTask выполнилась за приемлемое время
		return nil
	case <-timeoutCtx.Done(): // произошла отмена контекста по тайм-ауту
		return timeoutCtx.Err() // возвращаем ошибку, типа context deadline exceeded
	}
}

func executeTask() {
	time.Sleep(10 * time.Second)
}
