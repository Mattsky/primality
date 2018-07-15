package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func addr_check(address string) bool {

	targetIPPos := strings.Index(address, ":")

	if targetIPPos == -1 {
		fmt.Println("Invalid address entered.")
		return false
	}

	targetPortPos := strings.LastIndex(address, ":")

	if targetPortPos == -1 {
		fmt.Println("Invalid port entered.")
		return false
	}

	targetIP := address[0:targetIPPos]

	targetPort := address[targetPortPos+1:len(address)]

	//fmt.Println(targetIP)
	//fmt.Println(targetPort)

	// Check IP address is valid
	ipcheck := net.ParseIP(targetIP)
	//fmt.Println(ipcheck)
	
	if ipcheck != nil {
		
		// Check port is valid
		testport, err := strconv.Atoi(targetPort)

		// Check there was no error with conversion
		if err == nil {

			// If port is greater than 0...
			if testport >= 1 {

				// ... and less than 65535, then it's in the valid range
				if testport < 65535 {
					fmt.Println("Address appears OK. Proceeding.")
					return true
				} else {
					fmt.Println("Port is out of valid range.")
					return false
				}
			} else {
				fmt.Println("Negative ports aren't valid.")
				return false
			}
		} else {
			fmt.Println("Port doesn't seem valid.")
			return false
		}
	} else {
		fmt.Println("IP address seems invalid.")
		return false
	}
	return false
}

func input_check(text string) bool {
	// Is the input actually a number?
	_, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	
	if err == nil{    

		// Is the input greater than 0?
		testint, _ := strconv.Atoi(strings.TrimSuffix(text, "\n"))

		if testint > 0 { 

		    fmt.Println("Input valid.")
		    return true

	    // Negative number detected - reject it.
    	} else {
    		fmt.Print("Negative numbers won't work here. \n")
    		return false
    	}

    // General non-valid input caught. Reject it.
    } else {
    	fmt.Print("Invalid input detected. \n")
    	return false
    }
    return false

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please pass an address and port (e.g. 127.0.0.1:8000) to the program.")
		os.Exit(1)
	}

	targetAddress := os.Args[1]

	// Check input address is valid
	if addr_check(targetAddress) == true { 	

		// Connect to target server
		conn, err := net.Dial("tcp", targetAddress)
		if err != nil {
			fmt.Println("Error connecting to server.")
			os.Exit(1)
		}
		for { 

	    	// Read in input to be tested from STDIN.
	    	reader := bufio.NewReader(os.Stdin)
	    	fmt.Print("Text to send: ")
	    	text, _ := reader.ReadString('\n')

	    	//input_check(text, conn)
	    	if input_check(text) == true {
	    		// Send the data to the server...
		   		fmt.Fprintf(conn, text + "\n")

		    	// ... and wait for the reply
		    	message, _ := bufio.NewReader(conn).ReadString('\n')

		    	// So, is it a prime number?
		    	fmt.Print("Prime number: "+message)
	    	}
	  	}
	}
}