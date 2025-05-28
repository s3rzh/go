package main

import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5) // канал запросов которые нужно обслужить.
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests) // не забываем закрыть (чтобы for range ниже мог завешиться)

    limiter := time.Tick(200 * time.Millisecond) // Канал limiter будет получать значение каждые 200мс. Это то, что регулирует скорость в нашей схеме.

    for req := range requests {
        <-limiter //  Блокируя прием от канала limiter перед обслуживанием каждого запроса, мы ограничиваем себя одним запросом каждые 200 миллисекунд.
        fmt.Println("request", req, time.Now())
    }
}
