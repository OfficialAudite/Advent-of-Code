package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var location1 []int
var location2 []int
var sums []int
var sums2 []int

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

	sortLocations()
	calculateSums()
	calculateHowManyTimes()
	sum := addAllSums(sums)
	sum2 := addAllSums(sums2)

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	fmt.Printf("Part 1 sum: %d\n", sum)
	fmt.Printf("Part 2 sum: %d\n", sum2)
}

func processLine(line string) {
	split := strings.Split(line, "   ")

	location, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("Error parsing location:", err)
		return
	}

	location1 = append(location1, location)

	location, err = strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("Error parsing location:", err)
		return
	}

	location2 = append(location2, location)
}

func sortLocations(){
	sort.Ints(location1)
	sort.Ints(location2)
}

func calculateSums(){
	for i := 0; i < len(location1); i++ {
		loc1 := location1[i]
		loc2 := location2[i]

		if loc1 > loc2 {
			sums = append(sums, loc1-loc2)
		} else if loc1 < loc2 {
			sums = append(sums, loc2-loc1)
		}
	}
}

func addAllSums(sumlist []int) int {
	var sum int
	for i := 0; i < len(sumlist); i++ {
		sum += sumlist[i]
	}
	return sum
}

func calculateHowManyTimes(){
	for i := 0; i < len(location1); i++ {
		tempAmmount := 0
		for j := 0; j < len(location2); j++ {
			if location1[i] == location2[j] {
				tempAmmount++
			}
		}
		tempCalculation := location1[i] * tempAmmount
		sums2 = append(sums2, tempCalculation)
	}
}