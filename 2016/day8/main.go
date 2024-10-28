package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	screen := makeScreen(50, 6)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\w+`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, 1)
		action := matches[0]
		line = strings.Replace(line, action, "", 1)
		switch action {
		case "rect":
			rect(line, &screen)
		case "rotate":
			rotate(line, &screen)
		}

	}
	fmt.Println("Part 1:", countPixels(screen))
	fmt.Println("Part 2:")
	printScreen(screen)
}

func countPixels(screen [][]int) int {
	count := 0
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] > 0 {
				count++
			}
		}
	}
	return count
}

func copyScreen(screen [][]int) [][]int {
	newScreen := make([][]int, len(screen))
	for i := range screen {
		newScreen[i] = make([]int, len(screen[i]))
		copy(newScreen[i], screen[i])
	}
	return newScreen
}

func makeScreen(width, height int) [][]int {
	screen := make([][]int, height)
	for i := range screen {
		screen[i] = make([]int, width)
	}
	return screen
}

func rect(line string, screen *[][]int) {
	line = strings.Replace(line, " ", "", -1)
	parts := strings.Split(line, "x")

	width, _ := strconv.Atoi(parts[0])
	height, _ := strconv.Atoi(parts[1])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			(*screen)[i][j] = 1
		}
	}
}

func rotate(line string, screen *[][]int) {
	re := regexp.MustCompile(`(\w+)\s\w=(\d+)\sby\s(\d+)`)
	matches := re.FindStringSubmatch(line)

	column, _ := strconv.Atoi(matches[2])
	amount, _ := strconv.Atoi(matches[3])

	if matches[1] == "column" {
		rotateColumn(screen, column, amount)
	} else {
		rotateRow(screen, column, amount)
	}
}

func rotateColumn(screen *[][]int, column, amount int) {
	screenCopy := copyScreen(*screen)
	for i := len(screenCopy) - 1; i >= 0; i-- {
		if screenCopy[i][column] > 0 {
			next := (i + amount) % len(screenCopy)
			(*screen)[next][column] += 1
			(*screen)[i][column] -= 1
		}
	}

}

func rotateRow(screen *[][]int, row, amount int) {
	screenCopy := copyScreen(*screen)
	for i := len(screenCopy[row]) - 1; i >= 0; i-- {
		if screenCopy[row][i] > 0 {
			next := (i + amount) % len(screenCopy[row])
			(*screen)[row][i] -= 1
			(*screen)[row][next] += 1
		}
	}
}

func printScreen(screen [][]int) {
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
