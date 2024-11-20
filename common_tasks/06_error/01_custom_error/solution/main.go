package main

// реализовать функцию, которая возвращает ошибку, не используя импорт других библиотек

type myError struct {}

func (e *myError) Error() string {
	return "my error"
}

func handle() error {
	return &myError{}
}

func main() {
	err := handle()
	if err != nil {
		println(err.Error())
	}
}

// Создаем собственную структуру ошибки myError.
// Реализуем интерфейс error с помощью метода Error(), который возвращает string. Это необходимо, чтобы myError удовлетворял интерфейсу error, который требует удовлетворения единственного метода Error(), который возвращает string.
// Функция handle() имитирует работу обычной функции, которая возвращает ошибку. Она возвращает указатель на экземпляр myError. Поскольку myError реализует интерфейс error, функция handle() возвращает значение типа error.
// В main, происходит вызов функции handle() точно также, как мы обычно вызываем любую другую функцию с возвратом ошибки. После проверки на nil, мы распечатываем ошибку. В нашем случае, у нас всегда ошибка не равна nil, и мы распечатываем my error.