package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	readFile, _ := os.Open(os.Args[1])
	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	sum_p1 := 0
	sum_p2 := 0
	char := 0
	checkd := false
	for fileScanner.Scan() {
		line := fileScanner.Text()

		for _, c := range line {
			char++
			if c == '(' {
				sum_p1++
			} else {
				sum_p1--
			}
			if sum_p1 == -1 && !checkd {
				checkd = true
				sum_p2 = char
			}
		}
	}

	fmt.Println("Part 1:", sum_p1)
	fmt.Println("Part 2:", sum_p2)

	duration := time.Since(start)
	fmt.Println(duration)
}
