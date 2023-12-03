package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type point struct {
	number   string
	position []int
	row      int
}

type symbol struct {
	symbol   string
	position []int
	row      int
}

type rowMatches struct {
	row    int
	number int
}

var numbers []point
var symbols []symbol
var partNumbers_p1 []int
var partNumbers_p2 []int

func main() {
	start := time.Now()
	readFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	rePoint := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^\w\d\.]`)

	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		reNumbers := rePoint.FindAllStringIndex(line, -1)
		reSymbols := reSymbol.FindAllStringIndex(line, -1)

		for _, n := range reNumbers {
			start, end := n[0], n[1]
			tempNumber := line[start:end]
			numbers = append(numbers, point{
				number:   tempNumber,
				position: []int{start, end - 1},
				row:      count,
			})
		}

		for _, s := range reSymbols {
			start, end := s[0], s[1]
			tempSymbol := line[start:end]
			symbols = append(symbols, symbol{
				symbol:   tempSymbol,
				position: []int{start - 1, end},
				row:      count,
			})
		}

		count++
	}

	for _, n := range numbers {
		for _, s := range symbols {
			if isAdjacent(s, n) {
				a, err := strconv.Atoi(n.number)
				if err != nil {
					fmt.Println(err)
				} else {
					partNumbers_p1 = append(partNumbers_p1, a)
				}
			}
		}
	}

	for _, sym := range symbols {
		if sym.symbol == "*" {
			var adjacentNumbers []int
			for _, num := range numbers {
				if isAdjacent(sym, num) {
					n, err := strconv.Atoi(num.number)
					if err != nil {
						fmt.Println(err)
						continue
					}
					adjacentNumbers = append(adjacentNumbers, n)
				}
			}
			if len(adjacentNumbers) == 2 {
				gearRatio := adjacentNumbers[0] * adjacentNumbers[1]
				partNumbers_p2 = append(partNumbers_p2, gearRatio)
			}
		}
	}

	sum_p1 := 0
	for _, n := range partNumbers_p1 {
		sum_p1 += n
	}

	sum_p2 := 0
	for _, n := range partNumbers_p2 {
		sum_p2 += n
	}

	fmt.Println("Sum part1:", sum_p1)
	fmt.Println("Sum part2:", sum_p2)
	duration := time.Since(start)
	fmt.Println(duration)
}

func isAdjacent(sym symbol, num point) bool {
	if num.row == sym.row || num.row == sym.row+1 || num.row == sym.row-1 {
		if (num.position[0] >= sym.position[0] && num.position[0] <= sym.position[1]) ||
			(num.position[1] >= sym.position[0] && num.position[1] <= sym.position[1]) {
			return true
		}
	}
	return false
}
