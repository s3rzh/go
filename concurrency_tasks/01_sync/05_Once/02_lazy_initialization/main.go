package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	db *sql.DB
}

var (
	dbConn *DatabaseConnection
	once   sync.Once
)

func GetDatabaseConnection() (*DatabaseConnection, error) {
	var initError error
	once.Do(func() {
		fmt.Println("Initializing database connection...")
		db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
		if err != nil {
			initError = fmt.Errorf("failed to open database: %v", err)
			return
		}
		if err = db.Ping(); err != nil {
			initError = fmt.Errorf("failed to ping database: %v", err)
			return
		}
		dbConn = &DatabaseConnection{db: db}
	})
	if initError != nil {
		return nil, initError
	}
	return dbConn, nil
}

func main() {
	// Simulate multiple goroutines trying to get the database connection
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn, err := GetDatabaseConnection()
			if err != nil {
				log.Printf("Goroutine %d: Error getting connection: %v\n", id, err)
				return
			}
			log.Printf("Goroutine %d: Got connection %p\n", id, conn)
		}(i)
	}
	wg.Wait()
}

// для работы нужно поднять postgres в докере
