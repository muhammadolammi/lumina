package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Topic: Buffered Channels
	// Giving it a capacity of 10 means the sender can "burst" 10 items
	// before it has to wait for a receiver.
	logs := make(chan int, 10)

	// Here, we use an unbuffered channel to demonstrate blocking behavior.
	// logs := make(chan int, 0)

	// 1. Start 3 Workers
	for i := range 3 {
		wg.Add(1)
		go worker(i, logs, &wg)
	}

	// 2. The Producer (Main) sends 30 logs
	for i := range 30 {
		logs <- i
		fmt.Printf("Main: Sent log %d to channel\n", i)
	}

	// 3. Close the channel so workers know no more data is coming
	close(logs)

	wg.Wait()
	fmt.Println("Lumina: All logs processed successfully.")
}

// worker is our "Analyzer"
func worker(id int, logs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// The 'range' loop over a channel is the best way to drain it.
	// It automatically stops when the channel is closed.
	for logID := range logs {
		fmt.Printf("Worker %d: Processing log %d\n", id, logID)
		time.Sleep(50 * time.Millisecond) // Simulating work
	}
	fmt.Printf("Worker %d: Shutting down...\n", id)
}
