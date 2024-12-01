package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Start timer
	start := time.Now()

	// Check for input file
	if len(os.Args) < 2 {
		fmt.Println("Please specify an input file.")
		return
	}

	// Open the input file
	readFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		if err := readFile.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Scan the file
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		processLine(line) // Placeholder for your line processing logic
	}

	// Check for scanning errors
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print execution time
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
}

// processLine processes a single line of input
func processLine(line string) {
	fmt.Println(line)
}
