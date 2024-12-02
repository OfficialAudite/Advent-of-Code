package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Report struct {
	data []int
	safe bool
}

var reports []Report

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Please specify an input file.")
		return
	}

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

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		processLine(line)
	}

	for i := 0; i < len(reports); i++ {
		checkReport(i)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	safeReports := checkHowManySafeReports()
	fmt.Printf("Part 1 safe reports: %d\n", safeReports)
}

func processLine(line string) {
	split := strings.Split(line, " ")
	tempNumArray := make([]int, len(split))
	for i := 0; i < len(split); i++ {
		tempNum, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		tempNumArray[i] = tempNum
	}
	reports = append(reports, Report{data: tempNumArray, safe: false})
}

func checkReport(index int) {
	report := &reports[index]

	numbers := report.data
	dir := 0 // 0: undefined, 1: increasing, -1: decreasing
	report.safe = true

	for i := 1; i < len(numbers); i++ {
		num := numbers[i-1]
		nextNum := numbers[i]

		diff := nextNum - num

		if Abs(diff) < 1 || Abs(diff) > 3 {
			report.safe = false
			return
		}

		if dir == 0 {
			if diff > 0 {
				dir = 1
			} else if diff < 0 {
				dir = -1
			}
		} else {
			if (diff > 0 && dir == -1) || (diff < 0 && dir == 1) {
				report.safe = false
				return
			}
		}
	}
}

func checkHowManySafeReports() int {
	var safeReports int
	for i := 0; i < len(reports); i++ {
		if reports[i].safe {
			safeReports++
		}
	}
	return safeReports
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}