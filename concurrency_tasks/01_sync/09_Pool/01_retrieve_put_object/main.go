package main

import (
	"fmt"
	"sync"
)

type MyObject struct {
	// fields
}

func main() {
	var pool = &sync.Pool{
		New: func() any { // или старый алиас interface{}
			return &MyObject{}
		},
	}

	// Retrieve an object
	obj := pool.Get().(*MyObject)

	// Use the object
	// ...

	// Put the object back in the pool
	pool.Put(obj)

	// Retrieve the object again
	reusedObj := pool.Get().(*MyObject)

	fmt.Println(reusedObj == obj) // This will typically print true
}
