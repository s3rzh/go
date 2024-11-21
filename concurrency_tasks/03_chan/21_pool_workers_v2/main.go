package main

import (
	"fmt"
)

//  Написать WorkerPool с заданной функцией.

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- f(j)
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    multiplier := func(x int) int {
	    return x * 10
    }

    for w := 1; w <= 3; w++ {
        go worker(w,  multiplier, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    
    close(jobs)

    for i := 1; i <= numJobs; i++ {
        fmt.Println(<-results)
    }
}

//  Нам нужно разбить процессы на несколько горутин (3 по кол-ву воркеров) — при этом не создавать новую горутину каждый раз, а просто переиспользовать уже имеющиеся.
