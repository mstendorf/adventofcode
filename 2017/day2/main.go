package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1(file)
}

func part1(file *os.File) {

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\s+`)
		parts := re.Split(line, -1)
		// line = re.ReplaceAllString(line, " ")
		// parts := strings.Split(line, " ")
		sum += diff(parts)
	}
	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", part2(file))
}

func part2(file *os.File) int {
	file.Seek(0, 0)
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`\s+`)
		parts := re.Split(line, -1)
		// line = re.ReplaceAllString(line, " ")
		// parts := strings.Split(line, " ")
		sum += find_even_division(parts)
	}
	return sum
}

func find_even_division(parts []string) int {
	for i, part := range parts {
		int_part, _ := strconv.Atoi(part)
		for j, part2 := range parts {
			if i == j {
				continue
			}
			int_part2, _ := strconv.Atoi(part2)
			if IsEvenDivision(int_part, int_part2) {
				return int_part / int_part2
			} else if IsEvenDivision(int_part2, int_part) {
				return int_part2 / int_part
			}
		}
	}
	return 0
}

func diff(parts []string) int {
	min := 999999999999999
	max := 0
	for _, part := range parts {
		int_part, _ := strconv.Atoi(part)
		if int_part < min {
			min = int_part
		}
		if int_part > max {
			max = int_part

		}
	}
	return max - min
}

func IsEvenDivision(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false
}
