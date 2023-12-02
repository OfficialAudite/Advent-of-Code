package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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
}
