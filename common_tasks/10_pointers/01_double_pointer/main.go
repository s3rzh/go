package main

import "fmt"

// Что выведет программа и как её исправить?

type Person struct {
	name string
	age  int
}

func changeperson(person *Person) {
	person = &Person{
		name: "Alex",
		age:  30,
	}
}

func main() {
	person := &Person{"Ivan", 20}
	fmt.Println("person is", person)
	changeperson(person)
	fmt.Println("person is", person)
}
