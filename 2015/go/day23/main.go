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
	register  string
	argument  int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := get_instruction_set(file)
	part1(instructions)
	part2(instructions)

}

func part1(instructions []Instruction) {
	registers := make(map[string]int)
	fmt.Println(registers)
	registers = run_instructions(instructions, registers)
	fmt.Println("Part 1", registers)
}

func part2(instructions []Instruction) {
	registers := make(map[string]int)
	registers["a"] = 1
	registers = run_instructions(instructions, registers)
	fmt.Println("Part 2", registers)
}

func run_instructions(instructions []Instruction, registers map[string]int) map[string]int {
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		// fmt.Println(i, instruction)
		switch instruction.operation {
		case "hlf":
			registers[instruction.register] /= 2
		case "tpl":
			registers[instruction.register] *= 3
		case "inc":
			registers[instruction.register]++
		case "jmp":
			i += instruction.argument - 1
		case "jie":
			if registers[instruction.register]%2 == 0 {
				i += instruction.argument - 1
			}
		case "jio":
			if registers[instruction.register] == 1 {
				i += instruction.argument - 1
			}
		default:
			fmt.Println("Invalid instruction")
		}

		if i >= len(instructions) || i < 0 {
			fmt.Println("End of instructions")
			fmt.Println(registers["b"])
			break
		}
	}
	return registers
}

func get_instruction_set(file *os.File) []Instruction {

	instructions := make([]Instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instruction := map_to_instruction(line)
		// fmt.Println(instruction)
		instructions = append(instructions, instruction)
	}
	return instructions
}

func map_to_instruction(line string) Instruction {
	parts := strings.Split(line, ",")
	sub_parts := strings.Split(parts[0], " ")
	if len(parts) == 2 {
		trimmed := strings.TrimSpace(parts[1])
		argument, _ := strconv.Atoi(trimmed)
		instruction := Instruction{sub_parts[0], sub_parts[1], argument}
		return instruction
	} else if sub_parts[0] == "jmp" {
		arg, _ := strconv.Atoi(sub_parts[1])
		return Instruction{sub_parts[0], "", arg}
	}

	return Instruction{sub_parts[0], sub_parts[1], 0}
}
