package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) {
	lightGrid := createGrid()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "turn on") {
			turnOn(line, lightGrid)
		} else if strings.HasPrefix(line, "turn off") {
			turnOff(line, lightGrid)
		} else if strings.HasPrefix(line, "toggle") {
			toggle(line, lightGrid)
		}
	}

	count := 0
	for i := 0; i < len(lightGrid); i++ {
		for j := 0; j < len(lightGrid[i]); j++ {
			if lightGrid[i][j] == 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(file *os.File) {
	lightGrid := createGrid()
	// displayGrid(lightGrid)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "turn on") {
			turnOn2(line, lightGrid)
		} else if strings.HasPrefix(line, "turn off") {
			turnOff2(line, lightGrid)
		} else if strings.HasPrefix(line, "toggle") {
			toggle2(line, lightGrid)
		}
	}

	count := 0
	for i := 0; i < len(lightGrid); i++ {
		for j := 0; j < len(lightGrid[i]); j++ {
			count += lightGrid[i][j]
		}
	}
	fmt.Println(count)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	part1(file)
	file.Seek(0, 0)
	part2(file)
}

func createGrid() [][]int {
	grid := make([][]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func getCoords(line string) (int, int, int, int) {
	line = strings.Replace(line, " through ", ",", 1)
	coords := strings.Split(line, ",")
	x1, _ := strconv.Atoi(coords[0])
	y1, _ := strconv.Atoi(coords[1])
	x2, _ := strconv.Atoi(coords[2])
	y2, _ := strconv.Atoi(coords[3])

	return x1, y1, x2, y2
}

func turnOn(line string, lightGrid [][]int) {
	// turn on 887,9 through 959,629

	line = strings.Replace(line, "turn on ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = 1
		}
	}

}

func turnOn2(line string, lightGrid [][]int) {
	// turn on 887,9 through 959,629

	line = strings.Replace(line, "turn on ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j]++
		}
	}

}

func turnOff(line string, lightGrid [][]int) {
	line = strings.Replace(line, "turn off ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = 0
		}
	}
}

func turnOff2(line string, lightGrid [][]int) {
	line = strings.Replace(line, "turn off ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = max(0, lightGrid[i][j]-1)
		}
	}
}

func toggle(line string, lightGrid [][]int) {
	line = strings.Replace(line, "toggle ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if lightGrid[i][j] == 0 {
				lightGrid[i][j] = 1
			} else {
				lightGrid[i][j] = 0
			}
		}
	}
}
func toggle2(line string, lightGrid [][]int) {
	line = strings.Replace(line, "toggle ", "", 1)
	x1, y1, x2, y2 := getCoords(line)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] += 2
		}
	}
}
