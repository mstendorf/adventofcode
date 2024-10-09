package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])

		sum := 2*l*w + 2*w*h + 2*h*l
		slack := l * w
		if w*h < slack {
			slack = w * h
		}
		if h*l < slack {
			slack = h * l
		}
		total += sum + slack

		fmt.Println(scanner.Text(), "=", sum)
	}
	fmt.Println(total)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])

		sum := 2*l + 2*w + 2*h - 2*max(max(l, w), h) + l*w*h
		total += sum

		fmt.Println(scanner.Text(), "=", sum)
	}
	fmt.Println(total)
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	defer file.Close()

	part2(file)

}
