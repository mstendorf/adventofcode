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
	toggled   bool
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
		line := scanner.Text()
		parts := strings.Split(line, " ")
		instructions = append(instructions, Instruction{operation: parts[0], arguments: parts[1:]})
	}
	part1(instructions)

	part2(instructions)

}

func part1(instructions []Instruction) {
	registers := make(map[string]int)
	registers["a"] = 7
	execute(instructions, &registers)
	fmt.Println("Part 1", registers["a"])
}

func part2(instructions []Instruction) {
	// this should probably be optimized, but was faster to just let it crunch than to optimize
	registers := make(map[string]int)
	registers["a"] = 12
	for i := 0; i < len(instructions); i++ {
		instructions[i].toggled = false
	}
	execute(instructions, &registers)
	fmt.Println("Part 2", registers["a"])
}

func execute(instructions []Instruction, registers *map[string]int) {
	instructionsCpy := make([]Instruction, len(instructions))
	copy(instructionsCpy, instructions)

	for i := 0; i < len(instructionsCpy); i++ {
		instruction := instructionsCpy[i]

		switch instruction.operation {
		case "cpy":
			if instruction.toggled {
				instruction.toggled = false
				jmp := jnz(&instruction, registers)
				i += jmp
			} else {
				cpy(&instruction, registers)
			}
		case "tgl":
			tgl(&instruction, &instructionsCpy, registers, i)
		case "inc":
			inc(&instruction, registers)
		case "dec":
			dec(&instruction, registers)
		case "jnz":
			jmp := jnz(&instruction, registers)

			i += jmp
		default:
			panic("Unknown operation")
		}
	}
}

func tgl(instruction *Instruction, instructions *[]Instruction, registers *map[string]int, current_idx int) {
	if instruction.toggled {
		instruction.toggled = false
		inc(instruction, registers)
		return
	}
	register := instruction.arguments[0]
	distance := (*registers)[register]

	if current_idx+distance >= len(*instructions) {
		return
	}

	(*instructions)[current_idx+distance].toggled = !(*instructions)[current_idx+distance].toggled

}

func inc(instruction *Instruction, registers *map[string]int) {
	register := instruction.arguments[0]
	if instruction.toggled {
		instruction.toggled = false
		(*registers)[register]--
	} else {
		(*registers)[register]++
	}
}

func dec(instruction *Instruction, registers *map[string]int) {
	register := instruction.arguments[0]
	if instruction.toggled {
		instruction.toggled = false
		(*registers)[register]++
	} else {
		(*registers)[register]--
	}
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

	if instruction.toggled {
		instruction.toggled = false
		cpy(instruction, registers)
		return 0
	}

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
