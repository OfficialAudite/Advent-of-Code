package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type instruction struct {
	action string
	x1, y1 int
	x2, y2 int
}

var grid [1000][1000]bool
var grid2 [1000][1000]int

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Please specify an input file.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []instruction

	for scanner.Scan() {
		var inst instruction
		line := scanner.Text()
		parts := strings.Fields(line)

		if parts[0] == "toggle" {
			inst.action = parts[0]
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &inst.x1, &inst.y1, &inst.x2, &inst.y2)
		} else {
			inst.action = parts[0] + " " + parts[1]
			var action1, action2 string
			fmt.Sscanf(line, "%s %s %d,%d through %d,%d", &action1, &action2, &inst.x1, &inst.y1, &inst.x2, &inst.y2)
		}

		instructions = append(instructions, inst)
	}

	for _, inst := range instructions {
		for x := inst.x1; x <= inst.x2; x++ {
			for y := inst.y1; y <= inst.y2; y++ {
				switch inst.action {
				case "turn on":
					grid[x][y] = true
					grid2[x][y]++
				case "turn off":
					grid[x][y] = false
					if grid2[x][y] > 0 {
						grid2[x][y]--
					}
				case "toggle":
					grid[x][y] = !grid[x][y]
					grid2[x][y] += 2
				}
			}
		}
	}

	lightsOn := 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				lightsOn++
			}
		}
	}

	lightBrightness := 0
	for _, row := range grid2 {
		for _, light := range row {
			lightBrightness += light
		}
	}

	fmt.Println("Lights on:", lightsOn)
	fmt.Println("Light brightness:", lightBrightness)
	duration := time.Since(start)
	fmt.Println(duration)
}
