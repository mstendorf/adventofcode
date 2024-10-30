package main

import "fmt"

func main() {

	input := "^.^^^..^^...^.^..^^^^^.....^...^^^..^^^^.^^.^^^^^^^^.^^.^^^^...^^...^^^^.^.^..^^..^..^.^^.^.^......."

	grid := make([][]byte, 0)
	grid = append(grid, []byte(input))

	res := generate_grid(grid, 40)
	fmt.Println("Part 1", count_safe_tiles(res))
	res2 := generate_grid(grid, 400000)
	fmt.Println("Part 2", count_safe_tiles(res2))

}

func count_safe_tiles(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '.' {
				count++
			}
		}
	}
	return count
}

func generate_grid(grid [][]byte, length int) [][]byte {
	for i := 1; i < length; i++ {
		row := make([]byte, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			left := byte('.')
			center := grid[i-1][j]
			right := byte('.')
			if j > 0 {
				left = grid[i-1][j-1]
			}
			if j < len(grid[0])-1 {
				right = grid[i-1][j+1]
			}

			if (left == '^' && center == '^' && right == '.') ||
				(left == '.' && center == '^' && right == '^') ||
				(left == '^' && center == '.' && right == '.') ||
				(left == '.' && center == '.' && right == '^') {
				row[j] = '^'
			} else {
				row[j] = '.'
			}
		}
		grid = append(grid, row)
	}
	return grid
}
