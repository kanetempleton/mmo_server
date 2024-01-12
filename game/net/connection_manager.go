// Package net provides networking functionality for the server.
package net

import (
	"fmt"
	"net"
	"sync"
)

// ConnectionManager manages client connections.
type ConnectionManager struct {
	connections map[int]*Connection // Use connectionID as the key
	mu          sync.RWMutex        // Add a mutex for concurrent access.
}

// NewConnectionManager creates a new ConnectionManager instance.
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[int]*Connection),
	}
}


// HandleConnection handles a new incoming connection.
func (cm *ConnectionManager) HandleConnection(conn net.Conn) {
	connection := NewConnection(conn)

	// Add the connection to the manager's collection.
	cm.addConnection(connection)

	// Implement your logic for handling the connection.

	// Example: Read and process incoming data.
	buffer := make([]byte, 1024)
	for {
		n, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}

		// Process received data.
		cm.receiveData(connection, buffer[:n])
	}

	// Connection is closed, remove it from the manager.
	cm.removeConnection(connection)
}

// receiveData processes the received data.
func (cm *ConnectionManager) receiveData(connection *Connection, data []byte) {
	// Implement your logic for processing received data.
	// For example, print the received message along with the connection ID.
	fmt.Printf("Received from Connection ID %d: %s\n", connection.connectionID, data)

	//cm.SendMessageDirect(connection,"got your message")
}




// addConnection adds a connection to the manager.
func (cm *ConnectionManager) addConnection(connection *Connection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Check if the connectionID is already in use (unlikely due to recycling)
	if _, exists := cm.connections[connection.connectionID]; exists {
		fmt.Printf("Connection ID %d already exists.\n", connection.connectionID)
		return
	}

	cm.connections[connection.connectionID] = connection

	// Print a statement indicating that a client has connected.
	fmt.Printf("Client connected: Connection ID %d\n", connection.connectionID)
}


// removeConnection removes a connection from the manager.
func (cm *ConnectionManager) removeConnection(connection *Connection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Handle disconnection before removing the connection.
	cm.handleDisconnection(connection)

	// Remove the connection from the manager.
	delete(cm.connections, connection.connectionID)
}

// removeConnectionByID removes a connection from the manager by ID.
func (cm *ConnectionManager) RemoveConnectionByID(id int) {
	fmt.Printf("Connection ID %d being removed\n", id)
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Find the connection with the specified connectionID and remove it.
	if conn, exists := cm.connections[id]; exists {
		// Handle disconnection before removing the connection.
		cm.handleDisconnection(conn)
		

		// Remove the connection from the manager.
		delete(cm.connections, id)
		return
	}

	// Connection with the specified ID not found.
	fmt.Printf("Connection ID %d not found\n", id)
}



// handleDisconnection handles the cleanup when a connection is closed.
func (cm *ConnectionManager) handleDisconnection(connection *Connection) {
	// Implement your logic for handling disconnection.
	// For example, you can log the disconnection event.
	connection.Close()
	fmt.Printf("Client disconnected: Connection ID %d\n", connection.connectionID)
}

// GetConnections returns a slice of all active connections.
func (cm *ConnectionManager) GetConnections() []*Connection {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	connections := make([]*Connection, 0, len(cm.connections))
	for _, conn := range cm.connections {
		connections = append(connections, conn)
	}
	return connections
}

// SendMessage sends a message to a specific connection by ID.
func (cm *ConnectionManager) SendMessage(id int, message string) error {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	// Find the connection with the specified ID.
	if conn, exists := cm.connections[id]; exists {
		// Write the message to the connection.
		_, err := conn.Write([]byte(message))
		return err
	}

	// Connection with the specified ID not found.
	return fmt.Errorf("Connection ID %d not found", id)
}


// SendMessage sends a message to a specific connection.
func (cm *ConnectionManager) SendMessageDirect(conn *Connection, message string) error {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	// Check if the connection is in the manager's collection.
	if _, exists := cm.connections[conn.connectionID]; exists {
		// Write the message to the connection.
		_, err := conn.Write([]byte(message))
		fmt.Println("sent: "+message)
		return err
	}

	// Connection not found in the manager.
	return fmt.Errorf("Connection not found")
}

