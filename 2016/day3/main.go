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
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`^\s+(\d+)\s+(\d+)\s+(\d+)$`)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if isTriangle(matches[1:]) {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func isTriangle(triangle []string) bool {
	a, _ := strconv.Atoi(triangle[0])
	b, _ := strconv.Atoi(triangle[1])
	c, _ := strconv.Atoi(triangle[2])
	return a+b > c && a+c > b && b+c > a
}

func part2(file *os.File) {

	file.Seek(0, 0)
	// count := 0
	a := make([]string, 0, 3)
	b := make([]string, 0, 3)
	c := make([]string, 0, 3)
	re := regexp.MustCompile(`^\s+(\d+)\s+(\d+)\s+(\d+)$`)
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		for i, match := range matches[1:] {
			if i == 0 {
				a = append(a, match)
			} else if i == 1 {
				b = append(b, match)
			} else {
				c = append(c, match)
			}
		}

		if len(a) == 3 {
			if isTriangle(a) {
				count++
			}
			a = make([]string, 0, 3)
			if isTriangle(b) {
				count++
			}
			b = make([]string, 0, 3)
			if isTriangle(c) {
				count++
			}
			c = make([]string, 0, 3)
		}

	}

	fmt.Println("Part 2:", count)
}
