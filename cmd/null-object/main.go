package main

import (
	"fmt"
	nullObject "go-design-patterns/pkg/null-object"
)

// Application represents a system that uses a logger
type Application struct {
	logger nullObject.Logger
}

func NewApplication(logger nullObject.Logger) *Application {
	return &Application{logger: logger}
}

func (app *Application) Run() {
	app.logger.Info("Application is starting")
	// Simulate an error
	app.logger.Error("Something went wrong")
}

func main() {
	// Use the real logger
	realLogger := &nullObject.ConsoleLogger{}
	appWithRealLogger := NewApplication(realLogger)
	appWithRealLogger.Run()

	fmt.Println()

	// Use the null logger (no logging)
	nullLogger := &nullObject.NullLogger{}
	appWithNullLogger := NewApplication(nullLogger)
	appWithNullLogger.Run()
}
