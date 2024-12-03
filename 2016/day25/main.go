package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	arguments []string
}

func main() {

	// run the program, when it stops for a longer period on a given idx thats most likely the answer.
	// cba implementing loop detection
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]Instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		instructions = append(instructions, Instruction{operation: parts[0], arguments: parts[1:]})
	}

	part1(instructions)
}
func part1(instructions []Instruction) {
	registers := make(map[string]int)
	i := 1
	for {
		registers["a"] = i
		fmt.Println(i, registers["a"])

		execute(instructions, &registers)
		i++

	}
}

func execute(instructions []Instruction, registers *map[string]int) bool {
	instructionsCpy := make([]Instruction, len(instructions))
	copy(instructionsCpy, instructions)

	last := -1
	iterations := 0
	for i := 0; i < len(instructionsCpy); i++ {
		instruction := instructionsCpy[i]

		switch instruction.operation {
		case "cpy":
			cpy(&instruction, registers)
		case "inc":
			inc(&instruction, registers)
		case "dec":
			dec(&instruction, registers)
		case "jnz":
			jmp := jnz(&instruction, registers)

			i += jmp
		case "out":
			cur := out(instruction, registers)
			if cur == last {
				return false
			}
			last = cur

		default:
			panic("Unknown operation")
		}
		iterations++
	}
	return false
}

func out(instruction Instruction, registers *map[string]int) int {
	return (*registers)[instruction.arguments[0]]
}

func inc(instruction *Instruction, registers *map[string]int) {
	register := instruction.arguments[0]
	(*registers)[register]++
}

func dec(instruction *Instruction, registers *map[string]int) {
	register := instruction.arguments[0]
	(*registers)[register]--
}

func cpy(instruction *Instruction, registers *map[string]int) {
	if _, err := strconv.Atoi(instruction.arguments[1]); err == nil {
		// invalid argument, as the register variable is a number
		return
	}

	if val, err := strconv.Atoi(instruction.arguments[0]); err == nil {
		(*registers)[instruction.arguments[1]] = val
	} else {
		(*registers)[instruction.arguments[1]], _ = (*registers)[instruction.arguments[0]]
	}
}

func jnz(instruction *Instruction, registers *map[string]int) int {

	argument := -1
	if val, err := strconv.Atoi(instruction.arguments[0]); err == nil {
		// first var is integer
		argument = val
	} else {
		// first var is register
		argument = (*registers)[instruction.arguments[0]]
	}

	if argument != 0 {
		if val, err := strconv.Atoi(instruction.arguments[1]); err == nil {
			// second var is integer
			return val - 1
		} else {
			// second var is register
			return (*registers)[instruction.arguments[1]] - 1
		}

	}
	return 0
}
