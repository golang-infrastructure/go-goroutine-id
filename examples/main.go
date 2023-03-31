package main

import (
	"fmt"
	goroutine_id "github.com/golang-infrastructure/go-goroutine-id"
	"sync"
)

func main() {

	var goroutineId int
	var err error

	// Start a new coroutine and see what its id is
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineId, err = goroutine_id.GetGoroutineID()
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
	fmt.Println(fmt.Sprintf("Goroutine ID: %d", goroutineId))
	// Output:
	// Goroutine ID: 6

	// Look at the id of the main coroutine
	goroutineId, err = goroutine_id.GetGoroutineID()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Main Goroutine ID: %d", goroutineId))
	// Output:
	// Main Goroutine ID: 1

}
