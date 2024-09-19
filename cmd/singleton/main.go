package main

import "fmt"

func main() {
	// db := GetSingletonDatabase()
	// pop := db.GetPopulation("Seoul")
	// fmt.Println("Pop of seoul = ", pop)

	// Fetching the singleton instance
	singleton1 := GetInstance()
	fmt.Println(singleton1.value) // Output: This is a singleton instance

	// Fetching the singleton instance again to ensure it's the same instance
	singleton2 := GetInstance()
	fmt.Println(singleton2.value) // Output: This is a singleton instance

	// Check if both are the same instance
	if singleton1 == singleton2 {
		fmt.Println("Both are the same instance")
	}
}
