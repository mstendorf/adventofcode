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

	instructions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []string) {

	x := 1
	y := 1

	pincode := ""

	for _, ininstruction := range instructions {
		for _, instruction := range ininstruction {
			switch instruction {
			case 'U':
				if y > 0 {
					y--
				}
			case 'D':
				if y < 2 {
					y++
				}
			case 'L':
				if x > 0 {
					x--
				}
			case 'R':
				if x < 2 {
					x++
				}
			}
		}
		pincode += fmt.Sprintf("%d", y*3+x+1)
	}

	fmt.Println("Part 1:", pincode)
}

type Point struct {
	x int
	y int
}

func part2(instructions []string) {
	directions := map[rune]Point{
		'U': Point{0, -1},
		'D': Point{0, 1},
		'L': Point{-1, 0},
		'R': Point{1, 0},
	}
	keypad := [][]rune{
		{0, 0, '1', 0, 0},
		{0, '2', '3', '4', 0},
		{'5', '6', '7', '8', '9'},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}

	pos := Point{0, 2}
	pincode := ""
	for _, instruction := range instructions {
		for _, move := range instruction {
			newPos := Point{pos.x + directions[move].x, pos.y + directions[move].y}
			if newPos.x >= 0 && newPos.x <= 4 && newPos.y >= 0 && newPos.y <= 4 && keypad[newPos.y][newPos.x] != 0 {
				pos = newPos
			}
		}
		pincode += string(keypad[pos.y][pos.x])
	}

	fmt.Println("Part 2:", pincode)
}
