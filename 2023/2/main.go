package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	readFile, err := os.Open(`.\2023\2\input.txt`)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// currency
	red := 12
	green := 13
	blue := 14

	possibleGames := []string{}
	gamePower := []int{}

	for fileScanner.Scan() {
		possible := true
		r := 0
		g := 0
		b := 0
		line := fileScanner.Text()

		currentGame := strings.Replace(line[:strings.Index(line, ":")], "Game ", "", -1)
		line = line[strings.Index(line, ":")+2:]

		ongoingSets := strings.Split(line, "; ")

		for _, set := range ongoingSets {
			currentSets := strings.Split(set, ", ")
			for _, price := range currentSets {
				if strings.Contains(price, "red") {

					price = strings.Split(price, " ")[0]
					price, err := strconv.Atoi(price)

					if err != nil {
						fmt.Println(err)
					}

					if price > red {
						possible = false
					}

					if r < price {
						r = price
					}

				} else if strings.Contains(price, "green") {

					price = strings.Split(price, " ")[0]
					price, err := strconv.Atoi(price)

					if err != nil {
						fmt.Println(err)
					}

					if price > green {
						possible = false
					}

					if g < price {
						g = price
					}

				} else if strings.Contains(price, "blue") {
					price = strings.Split(price, " ")[0]
					price, err := strconv.Atoi(price)

					if err != nil {
						fmt.Println(err)
					}

					if price > blue {
						possible = false
					}

					if b < price {
						b = price
					}

				}
			}
		}

		if possible {
			possibleGames = append(possibleGames, currentGame)
		}

		gamePower = append(gamePower, r*g*b)
	}

	sum_p1 := 0
	for _, v := range possibleGames {
		num, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println(err)
		}

		sum_p1 += num
	}

	sum_p2 := 0
	for _, v := range gamePower {
		sum_p2 += v
	}

	fmt.Println("Current possible games part 1:", possibleGames)
	fmt.Println("Sum part 1:", sum_p1)
	fmt.Println("Sum part 2:", sum_p2)
	readFile.Close()
	duration := time.Since(start)
	fmt.Println(duration)
}
