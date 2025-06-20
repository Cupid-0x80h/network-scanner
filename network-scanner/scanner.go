package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// worker function to scan a single port.
// It takes a channel of ports to scan and a channel to report results.
func worker(ports chan int, results chan int, target string, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second) // 1-second timeout
		if err != nil {
			// Port is closed or filtered
			continue
		}
		conn.Close()
		results <- p // Send the open port number to the results channel
	}
}

func main() {
	var target string
	var choice int
	var startPort, endPort int

	// Get target IP from user
	fmt.Print("Enter the target IP address to scan: ")
	fmt.Scanln(&target)

	// Display menu and get user choice
	fmt.Println("\nSelect a scan option:")
	fmt.Println("1. Default Ports (1-1024)")
	fmt.Println("2. All Ports (1-65535)")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		startPort = 1
		endPort = 1024
	case 2:
		startPort = 1
		endPort = 65535
	default:
		fmt.Println("Invalid choice. Exiting.")
		return
	}

	fmt.Printf("\nScanning %s...\n", target)

	// --- Setup for Concurrent Scanning ---
	
	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to feed ports to the workers
	ports := make(chan int, 200) // Buffered channel for ports
	
	// Create a channel to collect results
	results := make(chan int)

	// Start 100 worker goroutines for concurrent scanning
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(ports, results, target, &wg)
	}

	// Send all ports to be scanned to the ports channel
	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
		close(ports) // Close the channel when all ports are sent
	}()

	var openPorts []int
	// Goroutine to collect results and close the results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read from the results channel until it is closed
	for port := range results {
		openPorts = append(openPorts, port)
	}

	// --- Display Results ---

	sort.Ints(openPorts) // Sort the ports for clean output

	if len(openPorts) == 0 {
		fmt.Println("\nNo open ports found.")
	} else {
		fmt.Println("\nOpen ports found:")
		// Convert integer slice to string slice for joining
		var portStrings []string
		for _, p := range openPorts {
			portStrings = append(portStrings, strconv.Itoa(p))
		}
		fmt.Println(strings.Join(portStrings, ", "))
	}
}