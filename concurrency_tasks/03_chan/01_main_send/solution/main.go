package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)
	c <- "John"
	fmt.Println("main() stopped")
}

// Разберем программу по шагам:

// Мы объявили функцию greet, которая принимает канал c как аргумент. В этой функции мы считываем данные из канала c и выводим в консоль.
// В функции main программа сначала выводит "main() started".
// Затем мы, используя make, создаем канал c с типом даных string.
// Помещаем канал с в функцию greet и запускаем функцию как горутину, используя ключевое слово go.
// Теперь у нас имеется две горутины main и greet, main по-прежнему остается активной.
// Помещаем данные в канал с и в этот момент main блокируется до тех пор, пока другая горутина (greet) не считает данные из канала c (Именно вычитает из канал, но может не успеть их распечатать к консоль тк main разблокируется и завериштся). Планировщик Go планирует запуск greet и выполняет описанное в первом пункте.
// После чего main снова становится активной и выводит в консоль "main() stopped".

// На деле вывод программы будет недетерминированный (неопределённый, undefined)

// иногда 
main() started
Hello John!
main() stopped

// а иногда и горутина будет не успевать напечатать данные (хоть и вычитывает их из канала)
main() started
main() stopped
