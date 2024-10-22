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

	part1(file)
	part2(file)
}

func printGrid(grid [100][100]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countNeighborLights(grid [][]bool, i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if i+x >= 0 && i+x < len(grid) && j+y >= 0 && j+y < len(grid[i]) {

				if grid[i+x][j+y] {
					count++
				}
			}
		}
	}
	return count
}

func part1(file *os.File) {
	grid := parseGrid(file)
	for i := 0; i < 100; i++ {
		grid = animate(grid)
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(file *os.File) {
	grid := parseGrid(file)
	grid[0][0] = true
	grid[0][len(grid[0])-1] = true
	grid[len(grid)-1][0] = true
	grid[len(grid)-1][len(grid[0])-1] = true
	for i := 0; i < 100; i++ {
		grid = animate(grid)
		grid[0][0] = true
		grid[0][len(grid[0])-1] = true
		grid[len(grid)-1][0] = true
		grid[len(grid)-1][len(grid[0])-1] = true
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func animate(grid [][]bool) [][]bool {
	orgGrid := make([][]bool, len(grid))
	copy(orgGrid, grid)
	for i := 0; i < len(grid); i++ {
		orgGrid[i] = make([]bool, len(grid[i]))
		copy(orgGrid[i], grid[i])
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			lights := countNeighborLights(orgGrid, i, j)
			if orgGrid[i][j] {
				if lights != 2 && lights != 3 {
					grid[i][j] = false
				}
			} else {
				if lights == 3 {
					grid[i][j] = true
				}
			}

		}
	}
	return grid
}

func parseGrid(file *os.File) [][]bool {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	grid := make([][]bool, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		grid = append(grid, make([]bool, len(line)))
		for j, char := range line {
			if char == '.' {
				grid[i][j] = false
			} else if char == '#' {
				grid[i][j] = true
			}
		}
	}
	return grid
}
