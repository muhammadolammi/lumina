package main

import (
	"sync"
)

// Global Stats
type Stats struct {
	mu            sync.Mutex
	logsProcessed int64
}

type LogSource interface {
	GetNextLog() (string, bool) // returns the log message and a 'success' bool
}

// 2. Implementation A: A Mock Stream (for testing)
type MockSource struct {
	currentCount int
	maxLogs      int
}

type LogEntry struct {
	// The tag `json:"level"` maps the JSON key "level" to this field
	Level   string `json:"level"`
	Message string `json:"msg,omitempty"` // omitempty: ignore if empty
	// Error   string `json:"error,omitempty"`
}

// 3. Implementation B: A File Source (placeholder for now)
type FileSource struct {
	FileName string
}

func (f *FileSource) GetNextLog() (string, bool) {
	// We'll implement actual file reading later
	return "file log data", true
}

// 4. Implementation C : A Error Source that logs out error
type ErrorSource struct {
	Messages chan string
}
