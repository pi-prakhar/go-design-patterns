package main

import "go-design-patterns/pkg/decorator"

func main() {
	// Create a concrete component
	component := &decorator.ConcreteComponent{}

	// Wrap it with a logging decorator
	decoratedComponent := &decorator.LoggingDecorator{Component: component}

	// Execute with the added logging functionality
	decoratedComponent.Execute()
}
