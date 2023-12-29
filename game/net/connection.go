// Package net provides networking functionality for the server.
package net

import (
	"fmt"
	"net"
)

// Connection represents a client connection.
type Connection struct {
	conn        net.Conn
	connectionID int // Unique identifier for each connection.
}

// counter to keep track of the next available connectionID
var connectionCounter = 0

// NewConnection creates a new Connection instance.
func NewConnection(conn net.Conn) *Connection {
	connection := &Connection{
		conn:        conn,
		connectionID: connectionCounter,
	}
	connectionCounter = (connectionCounter + 1) % 2000 // Recycle IDs if needed
	return connection
}

// Close closes the connection.
func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// Read reads data from the connection.
func (c *Connection) Read(buffer []byte) (int, error) {
	return c.conn.Read(buffer)
}

// Write writes data to the connection.
func (c *Connection) Write(data []byte) (int, error) {
	return c.conn.Write(data)
}

// PrintInfo prints information about the connection.
func (c *Connection) PrintInfo() {
	addr := c.conn.RemoteAddr()
	fmt.Printf("Connection Info:\n")
	fmt.Printf("  Connection ID: %d\n", c.connectionID)
	fmt.Printf("  Remote Address: %s\n", addr)
	// Add any other information you want to print.
	// For example, you might print the local address, connection status, etc.
}
