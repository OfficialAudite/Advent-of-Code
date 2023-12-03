package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// start timer
	start := time.Now()

	// read file
	readFile, err := os.Open(os.Args[1])

	// if error
	if err != nil {
		fmt.Println(err)
	}

	// scan file
	fileScanner := bufio.NewScanner(readFile)

	// close file
	defer readFile.Close()

	// split file
	fileScanner.Split(bufio.ScanLines)

	// for each line
	for fileScanner.Scan() {

		// read line
		line := fileScanner.Text()

		// print line
		fmt.Println(line)
	}

	// print time
	duration := time.Since(start)
	fmt.Println(duration)
}
