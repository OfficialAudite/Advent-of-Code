package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	Operator string
	Operand1 string
	Operand2 string
	Result   string
}

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
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	instructions := make(map[string]Instruction)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		instruction := parseInstruction(line)
		instructions[instruction.Result] = instruction
	}

	solved := make(map[string]uint16)
	fmt.Println("Part 1: ", solve(instructions, "a", solved))
	duration := time.Since(start)
	fmt.Println(duration)
}

func solve(instructions map[string]Instruction, wire string, solved map[string]uint16) uint16 {
	if val, exists := solved[wire]; exists {
		return val
	}

	instruction := instructions[wire]

	if instruction.Operator == "" {
		value, err := strconv.ParseUint(instruction.Operand1, 10, 16)
		if err != nil {
			value = uint64(solve(instructions, instruction.Operand1, solved))
		}
		solved[wire] = uint16(value)
		return uint16(value)
	}

	var operand1, operand2 uint16

	if instruction.Operand1 != "" {
		var err error
		val, err := strconv.ParseUint(instruction.Operand1, 10, 16)
		if err != nil {
			operand1 = solve(instructions, instruction.Operand1, solved)
		} else {
			operand1 = uint16(val)
		}
	}

	if instruction.Operand2 != "" {
		var err error
		val, err := strconv.ParseUint(instruction.Operand2, 10, 16)
		if err != nil {
			operand2 = solve(instructions, instruction.Operand2, solved)
		} else {
			operand2 = uint16(val)
		}
	}

	var result uint16
	switch instruction.Operator {
	case "AND":
		result = and(operand1, operand2)
	case "OR":
		result = or(operand1, operand2)
	case "NOT":
		result = not(operand1)
	case "LSHIFT":
		result = lshift(operand1, operand2)
	case "RSHIFT":
		result = rshift(operand1, operand2)
	}

	solved[wire] = result
	return result
}

func parseInstruction(line string) Instruction {
	words := strings.Fields(line)

	instruction := Instruction{}
	instruction.Result = words[len(words)-1]

	switch len(words) {
	case 3:
		instruction.Operand1 = words[0]
	case 4:
		instruction.Operator = words[0]
		instruction.Operand1 = words[1]
	case 5:
		instruction.Operand1 = words[0]
		instruction.Operator = words[1]
		instruction.Operand2 = words[2]
	}

	return instruction
}

func and(a, b uint16) uint16    { return a & b }
func or(a, b uint16) uint16     { return a | b }
func not(a uint16) uint16       { return ^a }
func lshift(a, b uint16) uint16 { return a << b }
func rshift(a, b uint16) uint16 { return a >> b }
