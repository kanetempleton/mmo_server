// Package net provides networking functionality for the server.
package net

import (
	"fmt"
	"net"
)

// Server represents the main server.
type Server struct {
	listener          net.Listener
	connectionManager *ConnectionManager
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		connectionManager: NewConnectionManager(),
	}
}

// Start starts the server and listens for incoming connections.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", ":43595")
	if err != nil {
		return err
	}
	s.listener = listener

	fmt.Println("Server is running on tcp://127.0.0.1:43595")

	go s.acceptConnections()

	return nil
}

// Stop stops the server and closes the listener.
func (s *Server) Stop() {
	if s.listener != nil {
		s.listener.Close()
	}
}

// acceptConnections handles incoming connections.
func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine.
		go func() {
			s.connectionManager.HandleConnection(conn)

			// Connection is closed, remove it from the manager.
			//s.connectionManager.removeConnection(NewConnection(conn))
		}()
	}
}

// get connection manager
func (s *Server) ConnectionManager() *ConnectionManager {
	return s.connectionManager
}
