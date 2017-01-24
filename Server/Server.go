package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Could not make a Connection over port 8080")
	}
	// defer listener.Close()

	// Loops infinitely for serving multiple connection requests
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting Connection")
		}
		fmt.Printf("CONNECTED!!\nConnection Address : %+v, %+v\n", connection.LocalAddr(), connection.RemoteAddr())

		go readFromConnection(connection)
	}
}

func readFromConnection(connection net.Conn) {
	var goodbyes bool
	reader := bufio.NewReader(connection)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			connection.Close()
		}

		start := time.Now()

		request := strings.ToUpper(string(line))

		var validKeyword bool
		if exists := strings.Contains(request, "HELLO"); exists {
			// Reply to Client "Hi"
			go respondClient(connection, "Hi !!\n")
			validKeyword = true
		}

		if exists := strings.Contains(request, "NAME"); exists {
			// Reply to Client "Chitty"
			go respondClient(connection, "My Name is Chitty.\n")
			validKeyword = true
		}

		if exists := strings.Contains(request, "BYE"); exists {
			// Reply to Client "Bye" or "Dont Go"
			if goodbyes {
				go respondClient(connection, "Bye!! Huh!\n")
			} else {
				go respondClient(connection, "Please! Don't Go! Let's Talk!\n")
				// goodbyes = !goodbyes
			}
			goodbyes = !goodbyes
			validKeyword = true
		}

		if !validKeyword {
			// For other inputs.
			go respondClient(connection, "Ummmmm......\n")
		}

		elapsed := time.Since(start)
		fmt.Printf("Time taken to process request from %+v : %s\n", connection.RemoteAddr(), elapsed.String())
	}
}

func respondClient(connection net.Conn, response string) {
	connection.Write([]byte(response))
}
