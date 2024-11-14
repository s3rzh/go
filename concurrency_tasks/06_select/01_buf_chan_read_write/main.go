package main

import "fmt"

func main() {
	bufChan := make(chan string, 1)
	bufChan <- "first"

  	// тк в буферезированом канале уже есть запись, то будет исполнен однозначно case str := <-bufChan: тк записать уже некуда (при готовности двух и более кэйсов - будет рандом, см пример ниже)
	select {
	case str := <-bufChan: // всегда этот вариант
		fmt.Println("read", str)
	case bufChan <- "second":
		fmt.Println("write", <-bufChan, bufChan)
	}

	bufChan2 := make(chan string, 2)
	bufChan2 <- "any"

  	// будет рандом тк есть что вычитывать из канала и есть ещё место в буфере
	select {
	case str2 := <-bufChan2:
		fmt.Println("read", str2)
	case bufChan2 <- "one":
		fmt.Println("write", <-bufChan2, <-bufChan2)
	}
}
