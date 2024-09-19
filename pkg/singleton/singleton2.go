package singleton

import (
	"sync"
)

// Singleton is the struct for which we will ensure only one instance exists
type Singleton struct {
	value string
}

// instance is the single instance of Singleton, initially nil
var instance *Singleton

// once is used to ensure the instance is created only once
var once sync.Once

// GetInstance returns the single instance of Singleton, creating it if necessary
func GetInstance() *Singleton {
	once.Do(func() {
		// This block runs only once
		instance = &Singleton{value: "This is a singleton instance"}
	})
	return instance
}
