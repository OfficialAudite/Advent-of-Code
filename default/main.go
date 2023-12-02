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
	readFile, err := os.Open(`.\1\input.txt`)

	// if error
	if err != nil {
		fmt.Println(err)
	}

	// scan file
	fileScanner := bufio.NewScanner(readFile)

	// split file
	fileScanner.Split(bufio.ScanLines)

	// for each line
	for fileScanner.Scan() {

		// read line
		line := fileScanner.Text()

		// print line
		fmt.Println(line)
	}

	// close file
	readFile.Close()

	// print time
	duration := time.Since(start)
	fmt.Println(duration)
}
