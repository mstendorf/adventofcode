package main

import (
	"fmt"
	"math"
)

func main() {
	input := 265149
	// input := 23

	// bottom right corner of the square is always an odd square
	// find the smallest odd square that is larger than the input
	i := 1
	for i*i < input {
		i += 2
	}

	// find the distance from the input to the nearest corner
	// the distance is the number of steps to the center
	corners := []int{i * i, i*i - i + 1, i*i - 2*i + 2, i*i - 3*i + 3}

	for _, corner := range corners {
		dist := math.Abs(float64(corner - input))
		if dist < float64(i-1/2) {
			fmt.Println("Part 1", i-1-int(dist))
			break
		}
	}

	part2(input)
}

func part2(input int) {

	coords := [][]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	x, y, dx := 0, 0, 0
	dy := -1
	spiral := make(map[string]int)
	spiral["0,0"] = 1

	for {
		total := 0
		for _, coord := range coords {
			key := fmt.Sprintf("%d,%d", x+coord[0], y+coord[1])
			if val, ok := spiral[key]; ok {
				total += val
			}
		}
		if total > input {
			fmt.Println("Part 2", total)
			break
		}
		if x == 0 && y == 0 {
			total = 1
		}
		spiral[fmt.Sprintf("%d,%d", x, y)] = total

		if (x == y) || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}
		x, y = x+dx, y+dy
	}
}
