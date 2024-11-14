package main

import "fmt"

func main() {
	bufChan := make(chan string, 1)
	bufChan <- "first"

  	// тк в буферезированом канале уже есть запись, то будет исполнен однозначно case str := <-bufChan: тк записать уже некуда (при готовности двух и более кэйсов - будет рандом)
	select {
	case str := <-bufChan:
		fmt.Println("read", str)
	case bufChan <- "second":
		fmt.Println("write", <-bufChan, bufChan)
	}
}
