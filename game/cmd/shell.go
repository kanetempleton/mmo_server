// Package cmd is the entry point for the server.
package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)

// Shell handles user input commands.
type Shell struct {
	launcher *Launcher
}

// NewShell creates a new Shell instance.
func NewShell(launcher *Launcher) *Shell {
	return &Shell{
		launcher: launcher,
	}
}

// Start starts listening for user input commands.
func (s *Shell) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		s.processCommand(command)
	}
}

// processCommand processes user input commands.
func (s *Shell) processCommand(command string) {
	args := strings.Fields(command)

	switch args[0] {
	case "connections":
		s.printConnectionsInfo()
	case "kick":
		if len(args) < 2 {
			fmt.Println("Usage: kick <ConnectionID>")
			return
		}

		connectionID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid Connection ID:", args[1])
			return
		}

		s.kickConnection(connectionID)
	case "message":
		if len(args) < 3 {
			fmt.Println("Usage: message <ConnectionID> <Message>")
			return
		}

		connectionID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid Connection ID:", args[1])
			return
		}

		message := strings.Join(args[2:], " ")
		//s.sendMessage(connectionID, message)
		packet := s.launcher.server.ConnectionManager().GetProtocol().SendMessagePacket(message)

		fmt.Printf("packet length is %d\n",packet.PayloadLength());

		// Send the message packet using SendPacket method
		err = s.launcher.server.ConnectionManager().SendPacket(connectionID, packet)

		//s.launcher.server.ConnectionManager().sendPacket();
	default:
		fmt.Println("Unknown command:", command)
	}
}



// printConnectionsInfo prints information about current connections.
func (s *Shell) printConnectionsInfo() {
	connections := s.launcher.server.ConnectionManager().GetConnections()

	fmt.Printf("Number of Connections: %d\n", len(connections))

	for _, conn := range connections {
		// You'll need to implement the printInfo method in the Connection type.
		// Assuming conn.printInfo() is a method you'll add to Connection.
		conn.PrintInfo()
	}
}

// kickConnection kicks a connection based on the specified Connection ID.
func (s *Shell) kickConnection(id int) {
	s.launcher.server.ConnectionManager().RemoveConnectionByID(id)
	fmt.Printf("Kicked Connection ID %d\n", id)
}


// sendMessage sends a message to a specific connection by ID.
func (s *Shell) sendMessage(connectionID int, message string) {
	err := s.launcher.server.ConnectionManager().SendMessage(connectionID, message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Message sent to Connection ID %d\n", connectionID)
	}
}