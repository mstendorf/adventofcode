package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	data := string(file)
	re := regexp.MustCompile(`\s+`)
	parts := re.Split(data, -1)

	blocks := make([]int, 0)
	for _, part := range parts {
		block, err := strconv.Atoi(part)
		// quick handling for whitespaces
		if err == nil {
			blocks = append(blocks, block)
		}
	}

	fmt.Println("Part 1", part1(blocks))
	fmt.Println("Part 2", part2(blocks))
}

func part1(blocksParam []int) int {
	blocks := make([]int, len(blocksParam))
	copy(blocks, blocksParam)

	seen := make(map[string]bool)
	seen[arrayToString(blocks)] = true

	cycles := 0
	for {
		idx := findBiggestIdx(blocks)
		blocksToDistribute := blocks[idx]
		blocks[idx] = 0

		for i := 1; i <= blocksToDistribute; i++ {
			blocks[(idx+i)%len(blocks)]++
		}

		if seen[arrayToString(blocks)] {
			break
		}
		seen[arrayToString(blocks)] = true
		cycles++
	}

	return cycles + 1
}

func part2(blocksParam []int) int {
	blocks := make([]int, len(blocksParam))
	copy(blocks, blocksParam)

	seen := make(map[string]int)
	seen[arrayToString(blocks)] = 0

	cycles := 0
	for {
		idx := findBiggestIdx(blocks)
		blocksToDistribute := blocks[idx]
		blocks[idx] = 0

		for i := 1; i <= blocksToDistribute; i++ {
			blocks[(idx+i)%len(blocks)]++
		}

		if cycle, ok := seen[arrayToString(blocks)]; ok {
			return cycles - cycle
		}
		seen[arrayToString(blocks)] = cycles
		cycles++
	}
}

func arrayToString(blocks []int) string {
	str := ""
	for _, block := range blocks {
		str += strconv.Itoa(block)
	}
	return str
}

func findBiggestIdx(blocks []int) int {
	biggest := 0
	idx := 0
	for i, block := range blocks {
		if block > biggest {
			biggest = block
			idx = i
		}
	}
	return idx
}
