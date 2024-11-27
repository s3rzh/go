package main

import (
	"fmt"
)

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true}
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)

}

// те итерируемся по мапе и добавляем новые ключи.
// резуьтат будет недетерминированным (непредсказуем)

// map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true]
// map[0:true 1:false 2:true 10:true 12:true 20:true]
// map[0:true 1:false 2:true 10:true 12:true 22:true 32:true 42:true 52:true 62:true]

// Вот что говорится в спецификации Go по поводу создания нового элемента мапы во время итераций:
// Если во время итерации создается элемент мапы, он может быть обработан во время итерации или пропущен. Выбор может варьироваться для каждого созданного элемента и от одной итерации к другой.
