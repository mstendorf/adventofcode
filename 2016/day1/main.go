package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ", ")

	part1(parts)
	part2(parts)
}

func part1(parts []string) {
	// Part 1
	start_x := 0
	start_y := 0
	x := start_x
	y := start_y

	// 0 = N, 1 = E, 2 = S, 3 = W
	direction := 0

	for _, part := range parts {
		// trim := strings.TrimSpace(part)
		if part[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction - 1) % 4
		}
		if direction == -1 {
			direction = 3
		}
		// fmt.Println("Direction:", direction)

		distance, err := strconv.Atoi(part[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 0:
			y += distance
		case 1:
			x += distance
		case 2:
			y -= distance
		case 3:
			x -= distance
		default:
			log.Fatal("Invalid direction")
		}
	}
	fmt.Println("Part 1:", math.Abs(float64(x))+math.Abs(float64(y)))
}

func part2(parts []string) {
	start_x := 0
	start_y := 0
	x := start_x
	y := start_y

	// 0 = N, 1 = E, 2 = S, 3 = W
	direction := 0
	visited := make(map[Point]bool)

	for _, part := range parts {
		// trim := strings.TrimSpace(part)
		if part[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction - 1) % 4
		}
		if direction == -1 {
			direction = 3
		}
		// fmt.Println("Direction:", direction)

		distance, err := strconv.Atoi(part[1:])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < distance; i++ {
			switch direction {
			case 0:
				y += 1
			case 1:
				x += 1
			case 2:
				y -= 1
			case 3:
				x -= 1
			default:
				log.Fatal("Invalid direction")
			}
			if visited[Point{x, y}] {
				fmt.Println("Visited twice:", x, y)
				fmt.Println("Part 2:", math.Abs(float64(x))+math.Abs(float64(y)))
				return
			}
			visited[Point{x, y}] = true
		}
	}
	fmt.Println("Part 2:", math.Abs(float64(x))+math.Abs(float64(y)))
}
