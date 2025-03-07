package main

import (
	"fmt"
	"sync"
)

// return channel for input numbers
func getInputChan() <-chan int {
	// make return channel
	input := make(chan int, 100)

	// sample numbers
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// run goroutine
	go func() {
		for _, num := range numbers {
			input <- num
		}
		// close channel once all numbers are sent to channel
		close(input)
	}()

	return input
}

// returns a channel which returns square of numbers
func getSquareChan(input <-chan int) <-chan int {
	// make return channel
	output := make(chan int, 100)

	// run goroutine
	go func() {
		// push squares until input channel closes
		for num := range input {
			output <- num * num
		}

		// close output channel once for loop finishes
		close(output)
	}()

	return output
}

// returns a merged channel of `outputsChan` channels
// this produce fan-in channel
// this is variadic function
func merge(outputsChan ...<-chan int) <-chan int {
	// create a WaitGroup
	var wg sync.WaitGroup

	// make return channel
	merged := make(chan int, 100)

	// increase counter to number of channels `len(outputsChan)`
	// as we will spawn number of goroutines equal to number of channels received to merge
	wg.Add(len(outputsChan))

	// function that accept a channel (which sends square numbers)
	// to push numbers to merged channel
	output := func(sc <-chan int) {
		// run until channel (square numbers sender) closes
		for sqr := range sc {
			merged <- sqr
		}
		// once channel (square numbers sender) closes,
		// call `Done` on `WaitGroup` to decrement counter
		wg.Done()
	}

	// run above `output` function as groutines, `n` number of times
	// where n is equal to number of channels received as argument the function
	// here we are using `for range` loop on `outputsChan` hence no need to manually tell `n`
	for _, optChan := range outputsChan {
		go output(optChan)
	}

	// run goroutine to close merged channel once done
	go func() {
		// wait until WaitGroup finishes
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// step 1: get input numbers channel
	// by calling `getInputChan` function, it runs a goroutine which sends number to returned channel
	chanInputNums := getInputChan() // тут уже происходит работа внутри getInputChan те заполняется выходной канал, можно посмотреть так time.Sleep(1 * time.Second) fmt.Println("chanInputNums len=", len(chanInputNums))

	// step 2: `fan-out` square operations to multiple goroutines
	// this can be done by calling `getSquareChan` function multiple times where individual function call returns a channel which sends square of numbers provided by `chanInputNums` channel
	// `getSquareChan` function runs goroutines internally where squaring operation is ran concurrently
	chanOptSqr1 := getSquareChan(chanInputNums) // начинает выполняется работа сразу после вызова
	chanOptSqr2 := getSquareChan(chanInputNums) // начинает выполняется работа сразу после вызова

	// step 3: fan-in (combine) `chanOptSqr1` and `chanOptSqr2` output to merged channel
	// this is achieved by calling `merge` function which takes multiple channels as arguments
	// and using `WaitGroup` and multiple goroutines to receive square number, we can send square numbers
	// to `merged` channel and close it
	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2) // начинает выполняется работа сразу после вызова

	// step 4: let's sum all the squares from 0 to 9 which should be about `285`
	// this is done by using `for range` loop on `chanMergedSqr`
	sqrSum := 0

	// run until `chanMergedSqr` or merged channel closes
	// that happens in `merge` function when all goroutines pushing to merged channel finishes
	// check line no. 86 and 87
	for num := range chanMergedSqr {
		sqrSum += num
	}

	// step 5: print sum when above `for loop` is done executing which is after `chanMergedSqr` channel closes
	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}

// Fan-in и Fan-out

// Fan-in — это стратегия мультиплексирования, при которой входы нескольких каналов объединяются
// в один выходной канал.
// Fan-out — это обратная операция, при которой один канал разделяется на несколько каналов.

// Пройдем по шагам.

// Получаем канал chanInputNums, посредством вызова функции getInputChan.
// Функция getInputChan создает канал и возвращает его как канал,
// доступный только для чтения, а также запускает анонимную горутину,
// которая последовательно помещает в канал числа из массива numbers и закрывает канал.

// Разделяем наш канал (fan-out) на два канала(chanOptSqr1 и chanOptSqr2),
// передавая его два раза функции getSquareChan.
// Функция getSquareChan создает канал и возвращает его как канал,
// доступный только для чтения, а также запускает анонимную горутину для
// вычисления квадрата чисел на основе данных канала, полученного в качестве аргумента функции.

// Собираем данные из каналов в один (fan-in), используя функцию merge.
// В функции merge мы создаем WaitGroup, а также новый канал(merged),
// где мы объединим все данные из списка каналов outputsChan, после,
// мы увеличиваем счетчик на основании числа полученных каналов,
// подготавливаем анонимную функцию для чтения данных из канала и
// группировки данных в наш новый канал merged, а также уменьшим
// значение счетчика, когда все данные из переданного канала будут считаны.
// Вызываем нашу анонимную функцию для каждого канала в качестве горутины.
// А так же создаем и стартуем еще одну анонимную горутину для того,
// чтобы дождаться выполнения операции объединения всех данных в один канал
// и после этого закрываем канал в рамках анонимной функции.
// После чего возвращаем наш новый канал merged.

// Считываем данные из канала chanMergedSqr используя for и range, и суммируем полученные данные.

// В конце выводим наш результат.
