package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	readFile, _ := os.Open(os.Args[1])
	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	sum_p1 := 0
	sum_p2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		w, h, l := 0, 0, 0
		fmt.Sscanf(line, "%dx%dx%d", &w, &h, &l)

		sum_p1 += 2*(w*h) + 2*(h*l) + 2*(w*l)
		sum_p1 += min(w*h, min(h*l, w*l))

		sum_p2 += 2*(w+h+l) - 2*max(w, max(h, l))
		sum_p2 += w * h * l
	}

	fmt.Println("Part 1:", sum_p1)
	fmt.Println("Part 2:", sum_p2)

	duration := time.Since(start)
	fmt.Println(duration)
}
