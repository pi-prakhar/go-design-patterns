package decorator

import (
	"fmt"
	"time"
)

// Component defines the behavior that can be decorated
type Component interface {
	Execute()
}

// ConcreteComponent is the original implementation of the Component interface
type ConcreteComponent struct{}

func (c *ConcreteComponent) Execute() {
	fmt.Println("Executing original behavior")
}

// LoggingDecorator wraps a component and adds logging
type LoggingDecorator struct {
	Component Component
}

func (d *LoggingDecorator) Execute() {
	start := time.Now()
	d.Component.Execute()
	fmt.Printf("Execution took %s\n", time.Since(start))
}
