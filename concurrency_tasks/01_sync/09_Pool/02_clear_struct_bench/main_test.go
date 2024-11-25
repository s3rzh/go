package main

import (
	"sync"
	"testing"
)

type Connection struct {
	ID   int
	Data []byte
}

var connPool = sync.Pool{
	New: func() interface{} {
		return &Connection{
			Data: make([]byte, 4096),
		}
	},
}

func handleRequest(conn *Connection) {
	// Обработка данных соединения
	for i := 0; i < len(conn.Data); i++ {
		conn.Data[i] = byte(i % 256)
	}
}

func resetConnection(conn *Connection) {
	conn.ID = 0
	for i := range conn.Data {
		conn.Data[i] = 0
	}
}

func BenchmarkWithSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn := connPool.Get().(*Connection)
		handleRequest(conn)
		resetConnection(conn) // Очистка структуры перед возвратом в пул
		connPool.Put(conn)
	}
}

func BenchmarkWithoutSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn := &Connection{
			Data: make([]byte, 4096),
		}
		handleRequest(conn)
	}
}

// go test -bench=. -benchmem

// В моём случае алокаций не случилось ни там ни там.. (Версия go 1.23.2)
// BenchmarkWithSyncPool-8           496591              2401 ns/op               0 B/op          0 allocs/op
// BenchmarkWithoutSyncPool-8        562040              2133 ns/op               0 B/op          0 allocs/op



