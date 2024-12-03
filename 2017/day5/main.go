package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	jumps := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		jump, _ := strconv.Atoi(line)
		jumps = append(jumps, jump)
	}
	fmt.Println("Part 1", part1(jumps))
	fmt.Println("Part 2", part2(jumps))
}

func part1(jumpsParam []int) int {
	jumps := make([]int, len(jumpsParam))
	copy(jumps, jumpsParam)

	steps := 0
	i := 0
	for i < len(jumps) {
		jump := jumps[i]
		jumps[i]++
		i += jump
		steps++
	}
	return steps
}

func part2(jumpsParam []int) int {
	jumps := make([]int, len(jumpsParam))
	copy(jumps, jumpsParam)

	steps := 0
	i := 0
	for i < len(jumps) {
		jump := jumps[i]
		if jump >= 3 {
			jumps[i]--
		} else {
			jumps[i]++
		}
		i += jump
		steps++
	}
	return steps
}
