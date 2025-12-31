package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
)

// worker is our "Analyzer"
func worker(id int, logs <-chan string, wg *sync.WaitGroup, stats *Stats) {
	defer wg.Done()
	for rawLog := range logs {
		logEntry := LogEntry{}

		err := json.Unmarshal([]byte(rawLog), &logEntry)
		if err != nil {
			fmt.Printf("Worker %d: ❌ Failed to parse JSON: %v\n", id, err)
			continue
		}
		// Now we can handle the log based on its data
		if logEntry.Level == "error" {
			fmt.Printf("Worker %d: 🚨 ALERT! Error detected: %s\n", id, logEntry.Message)
		} else {
			fmt.Printf("Worker %d: ✅ Processed %s: %s  \n", id, logEntry.Level, logEntry.Message)
		}

		// Increment our global counter safely
		atomic.AddInt64(&stats.logsProcessed, 1)

	}

}

func LogProcessor(log string, stats *Stats) {
	// process the log
	// simulate processing time
	fmt.Printf("LOGGER LOGGING: %s\n", log)
	// -- UPDATE THE PROCESSED LOGS COUNT SAFELY --//

	// using atomic operations to safely update shared stats , useful for simple int64 counters
	atomic.AddInt64(&stats.logsProcessed, 1)

	// using mutex lock to protect shared stats
	// stats.mu.Lock() // Grab the key
	// increment the processed logs count
	// stats.logsProcessed++
	// stats.mu.Unlock() // Release the key
}
