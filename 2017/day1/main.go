package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("Part 1", part1(line))
	fmt.Println("Part 2", part2(line))
}

func part1(line string) int {
	sum := 0
	if line[0] == line[len(line)-1] {
		sum += int(line[0] - '0')
	}
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			sum += int(line[i] - '0')
		}
	}
	return sum
}

func part2(line string) int {
	sum := 0
	for i := 0; i < len(line); i++ {
		if line[i] == line[(i+len(line)/2)%len(line)] {
			sum += int(line[i] - '0')
		}
	}
	return sum
}
