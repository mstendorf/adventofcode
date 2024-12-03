package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var grid [][]rune

	buildGrid(file, &grid)

	indexes := find_numbers(grid)
	numbers := make([]rune, 0)
	for _, index := range indexes {
		numbers = append(numbers, grid[index[1]][index[0]])
	}

	distances := make(map[int]map[int]int)
	for i, index := range indexes {
		distances[i] = make(map[int]int)
		for j, index2 := range indexes {
			if i != j {
				distances[i][j] = bfs_from_to(grid, index, index2)
			}
		}
	}
	pers := permutations(numbers)
	part1 := part1(distances, pers)
	fmt.Println("Part 1:", part1)
	part2 := part2(distances, pers)
	fmt.Println("Part 2:", part2)
}

func part1(distances map[int]map[int]int, permutations [][]rune) int {

	min := 1000000
	for _, per := range permutations {
		sum := 0
		for i := 0; i < len(per)-1; i++ {
			sum += distances[int(per[i]-'0')][int(per[i+1]-'0')]
		}
		if sum < min {
			min = sum
		}
	}
	return min
}

func part2(distances map[int]map[int]int, permutations [][]rune) int {

	min := 1000000
	for _, per := range permutations {
		sum := 0
		for i := 0; i < len(per)-1; i++ {
			sum += distances[int(per[i]-'0')][int(per[i+1]-'0')]
		}
		sum += distances[int(per[len(per)-1]-'0')][int(per[0]-'0')]
		if sum < min {
			min = sum
		}
	}
	return min
}

func permutations(arr []rune) [][]rune {
	if len(arr) == 1 {
		return [][]rune{arr}
	}
	perms := make([][]rune, 0)
	for i, val := range arr {
		rest := make([]rune, len(arr)-1)
		copy(rest, arr[:i])
		copy(rest[i:], arr[i+1:])
		for _, p := range permutations(rest) {
			perm := make([]rune, len(arr))
			perm[0] = val
			copy(perm[1:], p)
			perms = append(perms, perm)
		}
	}
	return perms
}

type Point struct {
	dst, x, y int
}

func bfs_from_to(grid [][]rune, fr []int, to []int) int {
	moves := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	q := make([]Point, 0)
	q = append(q, Point{0, fr[0], fr[1]})

	vis := make(map[Point]bool)
	point := Point{0, fr[0], fr[1]}
	vis[point] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.x == to[0] && cur.y == to[1] {
			return cur.dst
		}
		for _, move := range moves {
			ny, nx := cur.y+move[0], cur.x+move[1]
			if grid[ny][nx] != '#' && !vis[Point{0, nx, ny}] {
				q = append(q, Point{cur.dst + 1, nx, ny})
				vis[Point{0, nx, ny}] = true
			}
		}
	}
	return -1
}

// bfs distance from (y_fr,x_fr) to (y_to,x_to) assuming walls all around
// moves = set([(-1, 0), (1, 0), (0, 1), (0, -1)])
// def bfs_from_to(mp, fr, to):
//
//	q = deque([(0, fr)])
//	vis = set([fr])
//	while q:
//		dst, cur = q.pop()
//		if cur == to:
//			return dst
//		y, x = cur
//		for dy, dx in moves:
//			ny, nx = y + dy, x + dx
//			if mp[ny][nx] != '#' and (ny, nx) not in vis:
//				q.appendleft((dst+1, (ny, nx)))
//				vis.add((ny, nx))
//	return -1
func find_numbers(grid [][]rune) [][]int {
	indexes := make([][]int, 0)
	for y, row := range grid {
		for x, c := range row {
			if c >= '0' && c <= '9' {
				indexes = append(indexes, []int{x, y})
			}
		}
	}
	return indexes
}

func buildGrid(file *os.File, grid *[][]rune) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		*grid = append(*grid, row)
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}
