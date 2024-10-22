package main

import "fmt"

func main() {
	part1()
}

func part1() {
	x := 3010
	y := 3019
	size := 10000
	grid := make([][]int, size)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, size)
	}
	grid[0][0] = 20151125
	row := 0
	col := 0
	for {
		remainder := get_remainder(grid[row][col])
		if row == 0 {
			row = col + 1
			col = 0
		} else {
			row--
			col++
		}
		grid[row][col] = remainder
		// if row == 10 && col == 10 {
		// 	fmt.Println(grid[0][0])
		// 	fmt.Println(grid[0][1])
		// 	fmt.Println(grid[1][0])
		// 	fmt.Println(grid[1][1])
		// 	break
		// }
		// fmt.Println(row, col, grid[row][col])
		if row == x && col == y {
			fmt.Println(grid[x-1][y-1])
			break
		}
	}
}

func get_remainder(x int) int {
	// fmt.Println("remainder of ", x)
	return (x * 252533) % 33554393
}
