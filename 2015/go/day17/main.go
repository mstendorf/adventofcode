package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	containers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		container, _ := strconv.Atoi(line)
		containers = append(containers, container)
	}
	fmt.Println(part1(containers))
	fmt.Println(part2(containers))
}

func part1(containers []int) int {
	combos := Combinations(containers)
	return combos
}

func part2(containers []int) int {

	combos := MinCombinations(containers)
	return combos
}

func sum(containers []int) int {
	sum := 0
	for _, container := range containers {
		sum += container
	}
	return sum
}
func MinCombinations(set []int) int {
	length := uint(len(set))
	count := 0
	min := 1000

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		if sum(subset) == 150 {
			if len(subset) < min {
				min = len(subset)
				count = 1
			} else if len(subset) == min {
				count++
			}

		}

	}
	return count
}
func Combinations(set []int) int {
	length := uint(len(set))
	count := 0

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		if sum(subset) == 150 {
			count++
		}

	}
	return count
}
