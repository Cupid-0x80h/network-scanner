package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// --- Configuration ---
	var target_ip string
	if len(os.Args) > 1 {
		target_ip = os.Args[1]
	} else {
		target_ip = "scanme.nmap.org"
	}

	fmt.Printf("Scanning target: %s\n", target_ip)
	fmt.Println("==================================================")

	// --- The Scanner Logic in a Loop ---
	// Loop through the first 1024 ports
	for port := 1; port <= 1024; port++ {
		// We have to convert the integer port to a string for the address
		address := fmt.Sprintf("%s:%d", target_ip, port)
		// The logic is the same as before
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err == nil {
			// If no error, the port is open
			conn.Close()
			fmt.Printf("Port %d is OPEN\n", port)
		}
		// We do nothing if the port is closed (err is not nil)
	}    
	fmt.Println("==================================================")
	fmt.Println("Scan complete.")
}