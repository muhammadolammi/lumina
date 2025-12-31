package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Topic: Buffered Channels
	// Giving it a capacity of 10 means the sender can "burst" 10 items
	// before it has to wait for a receiver.

	logsC := make(chan string, 10)

	logSources := MockSource{
		currentCount: 0, maxLogs: 50,
	}
	// errorSource := ErrorSource{
	// 	Messages: make(chan string, 10),
	// }

	// Here, we use an unbuffered channel to demonstrate blocking behavior.
	// logs := make(chan int, 0)

	// lets create a stats struct to keep track of processed logs
	stats := &Stats{}

	// 1. Start 5 Workers
	for i := range 10 {
		wg.Add(1)
		go worker(i, logsC, &wg, stats)
	}

	// Simulate sending errors to the ErrorSource in a separate goroutine
	// go SendErrorToSource(50, &errorSource)

	produceLogs(&logSources, logsC)
	// produceLogs(&errorSource, logsC)
	close(logsC)

	// Wait for all workers to finish

	wg.Wait()
	fmt.Println("Lumina: All logs processed successfully. Total processed logs:", stats.logsProcessed)
}
