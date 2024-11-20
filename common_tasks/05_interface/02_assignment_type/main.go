package main

// Что распечатает?

func main() {
	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface == nil {
		println("It's nil!")
	}
}

// Ничего не распечается потому что в интерфейс != nil, в нем хранится информация о типе данных (struct{} - пустая структура)
// Если убрать присвоение iface = ptr, то интерфейс станет равным nil.
