package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
seed-to-soil map:
50 98 2
52 50 48
*/

type SeedMap struct {
	destination int
	source      int
	length      int
}

type Location struct {
	maps []SeedMap
}

type MegaSeed struct {
	start    int
	end      int
	newSeeds []int
}

var seeds []int

// [1]{destination: 2, source: 1, length: 1}
var seedMap []Location

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Please specify an input file.")
		return
	}
	readFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	lineCount := 0
	tempArray := []SeedMap{}
	megaSeeds := []MegaSeed{}

	totalLines := totalLines()

	seeds = []int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()

		seeds = append(seeds, getSeeds(lineCount, line)...)
		megaSeeds = append(megaSeeds, getMegaSeeds(lineCount, line)...)
		tempArray = append(tempArray, getMaps(lineCount, line)...)

		lineCount++
		if line == "" && lineCount > 2 || lineCount == totalLines {
			seedMap = append(seedMap, Location{maps: tempArray})
			tempArray = []SeedMap{}
		}
	}

	for i := range seeds {
		for _, location := range seedMap {
			for _, seedMap := range location.maps {
				destination := seedMap.destination
				source := seedMap.source
				length := seedMap.length

				if seeds[i] >= source && seeds[i] < source+length {
					var tempNum int
					if destination < source {
						tempNum = source - destination
						seeds[i] = seeds[i] - tempNum
					} else if destination > source {
						tempNum = destination - source
						seeds[i] = seeds[i] + tempNum
					}
					break
				}
			}
		}
	}

	for i, seed := range seeds {
		for i2, seed2 := range seeds {
			if i != i2 && seed < seed2 {
				seeds[i], seeds[i2] = seeds[i2], seeds[i]
			}
		}
	}

	sum_p1 := seeds[0]
	lowest := 0

	resultsChan := make(chan []int)

	var wg sync.WaitGroup
	for _, megaSeed := range megaSeeds {
		msStart := megaSeed.start
		msEnd := megaSeed.start + megaSeed.end

		proccesses := 10

		rangeSize := (msEnd - msStart) / proccesses

		for i := 0; i < proccesses; i++ {
			start := msStart + (i * rangeSize)
			end := start + rangeSize

			if i == proccesses-1 {
				end = msEnd
			}

			wg.Add(1)
			go processMegaSeedRange(&wg, start, end, seedMap, resultsChan)
		}

	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		for _, r := range result {
			if lowest == 0 || r < lowest {
				lowest = r
			}
		}
	}

	fmt.Println("Part 1:", sum_p1)
	fmt.Println("Part 2:", lowest)

	duration := time.Since(start)
	fmt.Println(duration)
}

func processMegaSeedRange(wg *sync.WaitGroup, start, end int, seedMap []Location, resultsChan chan []int) {
	defer wg.Done()
	tempNums := []int{}
	for i := start; i < end; i++ {
		seed := i
		for _, location := range seedMap {
			for _, sm := range location.maps {
				destination := sm.destination
				source := sm.source
				length := sm.length

				if seed >= source && seed < source+length {
					var tempNum int
					if destination < source {
						tempNum = source - destination
						seed = seed - tempNum
					} else if destination > source {
						tempNum = destination - source
						seed = seed + tempNum
					}
					tempNums = append(tempNums, seed)
					break
				}
			}
		}
	}
	defer func() { resultsChan <- tempNums }()
}

func findLowest(numbers []int) int {
	if len(numbers) == 0 {
		return 1337 // or some other default value
	}
	lowest := numbers[0]
	for _, n := range numbers {
		if n < lowest {
			lowest = n
		}
	}
	return lowest
}

func totalLines() int {
	readFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	totalLines := 0

	for fileScanner.Scan() {
		totalLines++
	}

	return totalLines
}

func getSeeds(lineCount int, line string) []int {
	seedReturn := []int{}
	if lineCount == 0 && line != "" && !strings.Contains(line, "map") {
		temp := strings.Split(line, ": ")
		tempSeeds := strings.Split(temp[1], " ")
		for _, seed := range tempSeeds {
			s, err := strconv.Atoi(seed)
			if err != nil {
				fmt.Println(err)
			}
			seedReturn = append(seedReturn, s)
		}
	}
	return seedReturn
}

func getMegaSeeds(lineCount int, line string) []MegaSeed {
	megaSeedReturn := []MegaSeed{}
	if lineCount == 0 && line != "" && !strings.Contains(line, "map") {
		temp := strings.Split(line, ": ")
		tempSeeds := strings.Split(temp[1], " ")
		for i := 0; i < len(tempSeeds)-1; i += 2 {
			nc1, err := strconv.Atoi(tempSeeds[i])
			nc2, err := strconv.Atoi(tempSeeds[i+1])

			if err != nil {
				fmt.Println(err)
			}
			megaSeedReturn = append(megaSeedReturn, MegaSeed{start: nc1, end: nc2})
		}
	}
	return megaSeedReturn
}

func getMaps(lineCount int, line string) []SeedMap {
	mapReturn := []SeedMap{}
	if lineCount > 0 && line != "" && !strings.Contains(line, "map") {
		temp := strings.Split(line, " ")
		num := []int{}
		for _, n := range temp {
			nc, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println(err)
			}
			num = append(num, nc)
		}
		mapReturn = append(mapReturn, SeedMap{destination: num[0], source: num[1], length: num[2]})
	}
	return mapReturn
}
