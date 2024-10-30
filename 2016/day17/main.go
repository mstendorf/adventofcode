package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Direction struct {
	x, y int
	char byte
}

var input = "vkjiggvb"

var shortest_path = ""
var longest_path = 0

func main() {
	walk("", 0, 0)
	fmt.Println("Part 1", shortest_path)
	fmt.Println("Part 2", longest_path)

}

func walk(route string, x, y int) {
	if x < 0 || x > 3 {
		return
	}
	if y < 0 || y > 3 {
		return
	}

	if x == 3 && y == 3 {
		if len(shortest_path) == 0 {
			shortest_path = route
		}
		if len(route) > longest_path {
			longest_path = len(route)
		}

		return
	}
	hash := GetMD5Hash([]byte(input + route))
	directions := getDirections(hash, x, y)
	for _, direction := range directions {
		walk(route+string(direction.char), x+direction.x, y+direction.y)
	}

}

func getDirections(hash string, x, y int) []Direction {
	directions := []Direction{}

	if hash[0] > 'a' {
		directions = append(directions, Direction{0, -1, 'U'})
	}
	if hash[1] > 'a' {
		directions = append(directions, Direction{0, 1, 'D'})
	}
	if hash[2] > 'a' {
		directions = append(directions, Direction{-1, 0, 'L'})
	}
	if hash[3] > 'a' {
		directions = append(directions, Direction{1, 0, 'R'})
	}
	return directions
}

func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
