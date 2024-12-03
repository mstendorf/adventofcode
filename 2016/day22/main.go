package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	size  int
	used  int
	avail int
	x, y  int
}

// var width = 3
// var height = 3

var width = 37
var height = 25

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	grid := buildGrid(file)
	part1(grid)

	part2(grid)
	printGrid(grid)
}

func part1(grid [][]Node) {

	count := countPair(grid, grid[0][0], 0, 0)
	fmt.Println("Part 1:", count)
}

func part2(grid [][]Node) {
	x, y := emptyNode(grid)
	// theese comments are relative to my presentation of the grid
	moves := 3 + y               // move over the wall and to the left of the grid
	moves += width - (x - 3) - 1 // move the empty node below the goal data
	moves += (width - 2) * 5     // move goal data to the 0,0 position

	fmt.Println("Part 2:", moves)

}

func emptyNode(grid [][]Node) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].used == 0 {
				return i, j
			}
		}
	}

	return -1, -1

}

func printGrid(grid [][]Node) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			node := grid[i][j]
			if node.used == 0 {
				fmt.Print("_")
			} else if hadAdjacentPair(grid, i, j) {
				fmt.Print("X")
			} else if node.size > 100 {
				fmt.Print("#")
			} else if i == width-1 && j == 0 {
				fmt.Print("G")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func hadAdjacentPair(grid [][]Node, x, y int) bool {
	node := grid[x][y]
	if x > 0 {
		if grid[x-1][y].avail >= node.used {
			return true
		}
	}
	if x < len(grid)-1 {
		if grid[x+1][y].avail >= node.used {
			return true
		}
	}
	if y > 0 {
		if grid[x][y-1].avail >= node.used {
			return true
		}
	}
	if y < len(grid[x])-1 {
		if grid[x][y+1].avail >= node.used {
			return true
		}
	}
	return false
}

func canMoveData(x, y int, grid [][]Node) bool {
	if grid[x][y].used == 0 {
		return false
	}
	return true
}

func countPair(grid [][]Node, current Node, x, y int) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			elem := grid[i][j]
			for k := 0; k < len(grid); k++ {
				for l := 0; l < len(grid[k]); l++ {
					if i == k && j == l {
						continue
					}
					if elem.used == 0 {
						continue
					}
					if elem.used <= grid[k][l].avail {
						count++
					}
				}
			}
		}
	}
	return count
}

func buildGrid(file *os.File) [][]Node {
	file.Seek(0, 0)
	grid := make([][]Node, width)
	for i := range grid {
		grid[i] = make([]Node, height)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	re := regexp.MustCompile(`x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)
	for scanner.Scan() {

		line := scanner.Text()
		parts := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])
		size, _ := strconv.Atoi(parts[3])
		used, _ := strconv.Atoi(parts[4])
		avail, _ := strconv.Atoi(parts[5])
		// fmt.Println(parts)
		node := Node{size: size, used: used, avail: avail, x: x, y: y}

		// fmt.Println(node, "x:", x, "y:", y)
		grid[x][y] = node
	}
	return grid
}
