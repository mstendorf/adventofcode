package main

import (
	"fmt"
	"slices"
)

type Position struct {
	x, y, steps int
}

type Visited struct {
	x, y int
}

func main() {
	visited := make([]Visited, 0)

	path := make([]Position, 0)
	path = append(path, Position{1, 1, 0})

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	part2 := 0
	found := false
outer:
	for i := 0; i < len(path); i++ {
		pos := path[i]
		if pos.steps > 50 && !found {
			found = true
			part2 = len(visited)
		}
		if !slices.Contains(visited, Visited{pos.x, pos.y}) {
			visited = append(visited, Visited{pos.x, pos.y})
		}
		for _, direction := range directions {
			newX, newY := pos.x+direction[0], pos.y+direction[1]
			if newX == 31 && newY == 39 {
				fmt.Println("Part 1", pos.steps+1)
				break outer
			}
			if !isWall(newX, newY) && !slices.Contains(visited, Visited{newX, newY}) {
				path = append(path, Position{newX, newY, pos.steps + 1})
			}
		}
	}
	fmt.Println("Part 2", part2)
}

func isWall(x, y int) bool {
	if x < 0 || y < 0 {
		return true
	}
	value := x*x + 3*x + 2*x*y + y + y*y + 1352
	return countBits(value)%2 == 1
}

func countBits(value int) int {
	count := 0
	for value > 0 {
		if value&1 == 1 {
			count++
		}
		value >>= 1
	}
	return count
}
