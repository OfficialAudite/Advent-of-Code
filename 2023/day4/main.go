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

type card struct {
	cardId         int
	winningNumbers []int
	numbersOnCard  []int
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
	cards := []card{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		winningNumbers := []int{}
		numbersOnCard := []int{}
		cardId := 0

		// card id
		fmt.Sscanf(line, "Card %d", &cardId)

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

		cards = append(cards, card{cardId, winningNumbers, numbersOnCard})
	}

	cardCounts := make(map[int]int)
	for _, c := range cards {
		cardCounts[c.cardId] = 1
	}

	for i := 0; i < len(cards); i++ {
		matches := 0
		for _, num := range cards[i].numbersOnCard {
			if contains(cards[i].winningNumbers, num) {
				matches++
			}
		}

		for j := 1; j <= matches; j++ {
			nextCardId := cards[i].cardId + j
			if nextCardId <= len(cards) {
				cardCounts[nextCardId] += cardCounts[cards[i].cardId]
			}
		}
	}

	totalCards := 0
	for _, count := range cardCounts {
		totalCards += count
	}

	fmt.Println("Total Points P1:", totalPoints)
	fmt.Println("Total Cards P2:", totalCards)
	duration := time.Since(start)
	fmt.Println("Duration:", duration)
}
