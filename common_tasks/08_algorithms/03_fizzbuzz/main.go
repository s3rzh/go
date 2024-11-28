package main

import "fmt"

// Написать программу, которая будет выводить числа от 1 до 100, при этом она должна выводить “Fizz”, если число кратно 3, “Buzz”, если число кратно 5, и “FizzBuzz”, если число кратно и 3, и 5 одновременно.

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { // i % 3 остаток от деления на 3, 1 % 3 = 1, 6 % 3 = 0, 7 % 3 = 1 итд 
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
