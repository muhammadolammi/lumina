package main

import (
	"fmt"
	"time"
)

func (m *MockSource) GetNextLog() (string, bool) {
	if m.currentCount >= m.maxLogs {
		return "", false
	}
	m.currentCount++
	return fmt.Sprintf(`{"level":"info", "msg":"Log #%d"}`, m.currentCount), true
}

func SendErrorToSource(numOfErrors int, errorSource *ErrorSource) {

	for i := range numOfErrors {
		errorSource.Messages <- fmt.Sprintf("error no %d", i)
		time.Sleep(1 * time.Second) // Simulate delay between errors
	}
	close(errorSource.Messages) // <--- THIS IS CRITICAL TO SIGNAL THE END OF ERRORS

}

func (e *ErrorSource) GetNextLog() (string, bool) {
	err, ok := <-e.Messages
	if !ok {
		// Channel was closed, signal the producer to stop
		return "", false
	}
	return fmt.Sprintf(`{"level":"error", "msg":"Log #%s"}`, err), true

}
