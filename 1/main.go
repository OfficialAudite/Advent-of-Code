package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func regexToNumber(s string) string {
	switch s {
	case "ne":
		s = "1"
	case "wo":
		s = "2"
	case "ree":
		s = "3"
	case "ur":
		s = "4"
	case "ive":
		s = "5"
	case "six":
		s = "6"
	case "ven":
		s = "7"
	case "ight":
		s = "8"
	case "nine":
		s = "9"
	}
	return s
}

func main() {
	start := time.Now()

	readFile, err := os.Open(`.\1\input.txt`)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()

	fileScanner.Split(bufio.ScanLines)

	count := 0
	re := regexp.MustCompile("(ne|wo|ree|ur|ive|six|ven|ight|nine|[1-9])")

	for fileScanner.Scan() {
		numbers := re.FindAllString(fileScanner.Text(), -1)

		for i, v := range numbers {
			numbers[i] = regexToNumber(v)
		}

		numberCombination, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])

		if err != nil {
			fmt.Println(err)
		}

		count += numberCombination
	}

	fmt.Println(count)

	duration := time.Since(start)
	fmt.Println(duration)
}
