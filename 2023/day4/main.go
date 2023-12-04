package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	readFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		winningNumbers := []int{}
		numbersOnCard := []int{}

		lineSplited := strings.Split(line, "|")
		cardNumbersBefore := strings.Fields(strings.Split(lineSplited[0], ":")[1])
		cardNumbersAfter := strings.Fields(lineSplited[1])

		for _, numberStr := range cardNumbersBefore {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			winningNumbers = append(winningNumbers, number)
		}

		for _, numberStr := range cardNumbersAfter {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			numbersOnCard = append(numbersOnCard, number)
		}

		// part 1
		points := 0
		for _, numberOnCard := range numbersOnCard {
			if contains(winningNumbers, numberOnCard) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}

	fmt.Println("Total Points P1:", totalPoints)
	duration := time.Since(start)
	fmt.Println("Duration:", duration)
}
