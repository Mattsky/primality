package main

import (
	"fmt"
	"math/big"
	"net"
	"bufio"
	"strconv"
)

func msg_handler(conn net.Conn) {

	fmt.Println("Connection established from", conn.RemoteAddr())

	// Close connection when this function ends or when EOF / dead conn detected
	defer func() {
		fmt.Println("Closing connection from", conn.RemoteAddr())
		conn.Close()
	}()
	
	// Infinite loop to keep connection open - not elegant, but I'm not a dev by trade (have mercy!)
	for {
		// Listen for message to process ending in newline
		message, _ := bufio.NewReader(conn).ReadString('\n')
	    
	    // DEBUG: show what was received
		fmt.Print("Message received from ", conn.RemoteAddr(), ": " + string(message))

	    // Run check on received message
		resultmsg := strconv.FormatBool(prime_check(message))
		
		// Is the number prime? Yes / no for debug.
		fmt.Print("PRIME: " + resultmsg + "\n")

		// Return result to client
		_, err := conn.Write([]byte(resultmsg + "\n"))

		// Check if connection is dead; if yes, end the loop & subroutine
		if err != nil {
			return
		}

	}

} 

func prime_check(message string) bool {
    	
	// Convert message string to big ol' int
	n := new(big.Int)
	n.SetString(message, 10)

	boolResult := n.ProbablyPrime(1)
	return boolResult
}

func main() {
	fmt.Println("Launching server...")

	// Listen on all available interfaces
	ln, _ := net.Listen("tcp", ":8000")

	// Main loop - run forever (or until program exits)
	for {
		
		// Accept connections
		conn, _ := ln.Accept()

		// Run goroutine to handle individual connections
		go msg_handler(conn)

	}
}