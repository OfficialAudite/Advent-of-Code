package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkForDouble(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func containsNonOverlappingPair(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		pair := s[i : i+2]
		restOfString := s[i+2:]
		if strings.Contains(restOfString, pair) {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	readFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	reVowels := regexp.MustCompile(`[aeiou]`)
	reDouble := regexp.MustCompile(`(aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz)`)
	reExclude := regexp.MustCompile(`ab|cd|pq|xy`)

	nice_p1 := 0
	nice_p2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		vowels := reVowels.FindAllString(line, -1)
		double := reDouble.FindAllString(line, -1)
		exclude := reExclude.FindAllString(line, -1)

		if len(vowels) >= 3 && len(double) >= 1 && len(exclude) == 0 {
			nice_p1++
		}

		if containsNonOverlappingPair(line) && checkForDouble(line) {
			nice_p2++
		}
	}

	fmt.Println(nice_p1)
	fmt.Println(nice_p2)
	duration := time.Since(start)
	fmt.Println(duration)
}
