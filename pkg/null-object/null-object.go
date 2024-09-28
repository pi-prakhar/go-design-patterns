package nullObject

import "fmt"

// Logger is the interface for logging
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// ConsoleLogger is a real logger that logs messages to the console
type ConsoleLogger struct{}

func (l *ConsoleLogger) Info(msg string) {
	fmt.Println("INFO:", msg)
}

func (l *ConsoleLogger) Error(msg string) {
	fmt.Println("ERROR:", msg)
}

// NullLogger is the null object that does nothing
type NullLogger struct{}

func (l *NullLogger) Info(msg string) {
	// Do nothing
}

func (l *NullLogger) Error(msg string) {
	// Do nothing
}
