package objectpool

import "sync"

// Connection is a simple struct that represents the object we want to pool
type Connection struct {
	ID int
}

// ConnectionPool manages a pool of connections
type ConnectionPool struct {
	pool  chan *Connection
	mu    sync.Mutex
	count int
}

// NewConnectionPool initializes a new pool of connections
func NewConnectionPool(size int) *ConnectionPool {
	return &ConnectionPool{
		pool: make(chan *Connection, size),
	}
}

// Get retrieves a connection from the pool or creates a new one if necessary
func (p *ConnectionPool) Get() *Connection {
	select {
	case conn := <-p.pool:
		return conn
	default:
		// Create a new connection if pool is empty
		p.mu.Lock()
		defer p.mu.Unlock()
		p.count++
		return &Connection{ID: p.count}
	}
}

// Put returns a connection back to the pool
func (p *ConnectionPool) Put(conn *Connection) {
	select {
	case p.pool <- conn:
		// Connection returned to pool
	default:
		// Pool is full, discard connection
	}
}
