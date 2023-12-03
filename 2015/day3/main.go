package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Point struct {
	X, Y int
}

func main() {
	start := time.Now()

	readFile, _ := os.Open(os.Args[1])
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	houses := make(map[Point]int)

	var santaX, santaY int
	var roboSantaX, roboSantaY int

	houses[Point{santaX, santaY}] = 2

	turn := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		for _, c := range line {
			if turn == 0 {
				switch c {
				case '^':
					santaY++
				case 'v':
					santaY--
				case '>':
					santaX++
				case '<':
					santaX--
				}
				houses[Point{santaX, santaY}]++
			} else {
				switch c {
				case '^':
					roboSantaY++
				case 'v':
					roboSantaY--
				case '>':
					roboSantaX++
				case '<':
					roboSantaX--
				}
				houses[Point{roboSantaX, roboSantaY}]++
			}
			turn = 1 - turn
		}
	}

	count := len(houses)
	fmt.Println("Number of houses that receive at least one present:", count)
	fmt.Println("Execution Time:", time.Since(start))
}
