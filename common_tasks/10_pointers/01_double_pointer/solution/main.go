package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeperson(person **Person) {
	*person = &Person{ // разыменовываем и переприсваиваем новый указтель
		name: "Alex",
		age:  30,
	}
}

func main() {
	person := &Person{"Ivan", 20} // переменная типа указатель
	fmt.Println("person is", person)
	changeperson(&person) // берём адрес переменной типа укатель
	fmt.Println("person is", person)
}

// Можно взять указатель на указатель

// Ещё можно решить через копирование, типа changeperson(person) ... func changeperson(person *Person)...  	*person = Person{... но это копирование обьекта
