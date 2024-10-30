package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	blocked := make([][2]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		block := [2]int{low, high}
		blocked = append(blocked, block)
	}
	first_unblocked := 0
	count := 0

	for i := 0; i < 4294967295; i++ {
		block := false
		for _, blocked_range := range blocked {
			if i >= blocked_range[0] && i <= blocked_range[1] {
				i = blocked_range[1]
				block = true
				break
			}
		}
		if !block {
			if first_unblocked == 0 {
				first_unblocked = i
			}
			count++
		}
	}
	fmt.Println("Part 1:", first_unblocked)
	fmt.Println("Part 2:", count)
}

func isBlocked(i int, blocked *[][2]int) bool {
	for _, block := range *blocked {
		if i >= block[0] && i <= block[1] {
			return true
		}
	}
	return false
}
