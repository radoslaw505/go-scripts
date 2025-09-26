package main

import (
	"fmt"
	"sync"
	"time"
)

/*
once_example.go

- Demonstrates sync.Once to perform one-time initialization concurrently.
- Also shows typical pattern to capture initialization error by storing it in a package-level variable.
*/

var initOnce sync.Once
var dbConn string // pretend resource handle
var initErr error

func initialize() {
	// long-running init
	time.Sleep(200 * time.Millisecond)
	// pretend success
	dbConn = "db-connection"
	// if error occurred: set initErr = errors.New("failed")
	fmt.Println("initialized")
}

func GetConnection() (string, error) {
	initOnce.Do(initialize)
	return dbConn, initErr
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn, err := GetConnection()
			fmt.Printf("goroutine %d got conn=%v err=%v\n", id, conn, err)
		}(i)
	}
	wg.Wait()
}
