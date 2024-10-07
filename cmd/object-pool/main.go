package main

import (
	"fmt"
	objectpool "go-design-patterns/pkg/object-pool"
)

func main() {
	// Create a pool with a max size of 2 connections
	pool := objectpool.NewConnectionPool(2)

	// Simulate requesting connections from the pool
	conn1 := pool.Get()
	fmt.Printf("Got connection %d\n", conn1.ID)

	conn2 := pool.Get()
	fmt.Printf("Got connection %d\n", conn2.ID)

	// Return the connections to the pool
	pool.Put(conn1)
	pool.Put(conn2)

	// Reuse the connection
	conn3 := pool.Get()
	fmt.Printf("Reused connection %d\n", conn3.ID)

	// Create a new connection since the pool was empty after reuse
	conn4 := pool.Get()
	fmt.Printf("Got new connection %d\n", conn4.ID)
}
