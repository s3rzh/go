package main

import (
	"fmt"
	"sync"
	"time"
)

var m = map[string]int{"a": 1}
var mtx = sync.RWMutex{}

func main() {
	go read()
	time.Sleep(1 * time.Second)
	go write()
	time.Sleep(1 * time.Second)
}

func read() {
	for {
		mtx.RLock()
		defer mtx.RUnlock()
		fmt.Println(m["a"])
	}
}

func write() {
	for {
		mtx.Lock()
		defer mtx.Unlock()
		m["a"] = 2
	}
}

// concurrent map read and map write

// чтобы отловить проблему с конкурентной записью
// можно запустить код с флагом data race

// Что будет, если в программе мы будем читать
// по ключу "a", а писать по ключу "b"?

// Все равно программа будет падать,
// потому что мы рассматриваем мапу как цельный объект

// Проблема решается добавлением мьютекса

// Чем отличаются Mutex и RWMutex?
// Обычный Mutex полносьтю лочит ресурсы и на чтение и на запись.
// Мы может безопасо читать из раздлеляемых ресурсов, а писать не можем.
// Когда мы используем RWMutex, горутины, которые приходят читать
// Как только у нас кто-то придет и захочет записать, мы дождемся,
// пока все горутины закончат чтение и только потом возьмем лок на запись
// если в этот момент кто-то захочет прочитать, ему придется подождать пока
// лок на запись будет снят и только тогда мы сможем прочитать