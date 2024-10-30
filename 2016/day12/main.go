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
	operation         string
	argument          int
	argument_register string
	register          string
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]Instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Do something with the line
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			instructions = append(instructions, Instruction{operation: parts[0], register: parts[1]})
		} else {
			// we have a jnz or cpy instruction
			arg_int, err := strconv.Atoi(parts[1])

			if err != nil {
				// we have a register
				instructions = append(instructions, Instruction{operation: parts[0], argument: -1, argument_register: parts[1], register: parts[2]})
			} else {
				// we have an integer
				instructions = append(instructions, Instruction{operation: parts[0], argument: arg_int, argument_register: "", register: parts[2]})

			}
		}
	}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []Instruction) {
	registers := make(map[string]int)
	execute(instructions, &registers)
	fmt.Println("Part 1", registers["a"])
}

func part2(instructions []Instruction) {
	registers := make(map[string]int)
	registers["c"] = 1
	execute(instructions, &registers)
	fmt.Println("Part 2", registers["a"])
}
func execute(instructions []Instruction, registers *map[string]int) {
	// registers := make(map[string]int)

	for i := 0; i < len(instructions); i++ {
		// fmt.Println(instructions[i])
		instruction := instructions[i]

		switch instruction.operation {
		case "cpy":
			cpy(instruction, registers)
		case "inc":
			(*registers)[instruction.register]++
		case "dec":
			(*registers)[instruction.register]--
		case "jnz":
			jmp := jnz(instruction, registers)
			i += jmp

		default:
			panic("Unknown operation")
		}
	}
}

func cpy(instruction Instruction, registers *map[string]int) {
	if instruction.argument != -1 {
		(*registers)[instruction.register] = instruction.argument
	} else {
		(*registers)[instruction.register] = (*registers)[instruction.argument_register]
	}
}

func jnz(instruction Instruction, registers *map[string]int) int {

	if instruction.argument != -1 {
		if instruction.argument != 0 {
			// jump
			jump, _ := strconv.Atoi(instruction.register)
			return jump - 1
		}
	} else {
		if (*registers)[instruction.argument_register] != 0 {
			// jump
			jump, _ := strconv.Atoi(instruction.register)
			return jump - 1
		}
	}
	return 0
}
