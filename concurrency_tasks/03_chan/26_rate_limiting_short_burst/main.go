package main

import (
    "fmt"
    "time"
)

func main() {
    burstyLimiter := make(chan time.Time, 3) // Мы можем разрешить короткие всплески запросов при сохранении общего ограничения скорости. Этот канал burstyLimiter будет позволять делать до 3 событий.

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now() // Заполняем канал, чтобы предоставить возможность ускорить.
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) { // Каждые 200мс мы будем пытаться добавлять новое значение в burstyLimiter, до своего предела в 3 значения.
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5) // Теперь смоделируем 5 входящих запросов. Первые 3 из них получат выгоду от вместимости burstyLimiter.
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
  
    for req := range burstyRequests {  // Мы обслуживаем первые 3 запроса сразу из-за использования ограничения скорости, затем обслуживаем оставшиеся 2 с задержками ~200мс каждый.
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
