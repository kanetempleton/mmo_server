// Package cmd is the entry point for the server.
package cmd

import (
	"fmt"
	"mmo_server/game/net"
	"os"
	"os/signal"
	"sync"
)

// Launcher manages the central coordination of the application.
type Launcher struct {
	server   *net.Server
	shell    *Shell
	shutdown chan struct{}
}

// NewLauncher creates a new Launcher instance.
func NewLauncher() *Launcher {
	launcher := &Launcher{
		server:   net.NewServer(),
		shutdown: make(chan struct{}),
	}

	// Pass the reference to the launcher instance to NewShell.
	launcher.shell = NewShell(launcher)

	return launcher
}

// Start starts the application components concurrently.
func (l *Launcher) Start() {
	var wg sync.WaitGroup

	// Start the server in a separate goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := l.server.Start(); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	// Start the shell in a separate goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.shell.Start()
	}()

	// Wait for all components to finish or a shutdown signal.
	select {
	case <-l.shutdown:
		fmt.Println("Received shutdown signal. Shutting down...")
	}

	// Stop the application components.
	l.Stop()

	// Wait for all components to finish.
	wg.Wait()
}

// Stop stops the application components.
func (l *Launcher) Stop() {
	// Stop the server.
	l.server.Stop()

	// Stop the shell.
	close(l.shutdown)
}

// ConnectionManager returns the connection manager from the server.
func (l *Launcher) ConnectionManager() *net.ConnectionManager {
	return l.server.ConnectionManager()
}

// waitForInterrupt waits for an interrupt signal to gracefully stop the application.
func waitForInterrupt(l *Launcher) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	<-sigCh

	// Signal the launcher to stop.
	close(l.shutdown)
}
