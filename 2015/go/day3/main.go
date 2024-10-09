package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type House struct {
	row int
	col int
}

type Santa struct {
	row int
	col int
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanRunes)

	seen := map[House]bool{}
	santa := Santa{0, 0}
	roboSanta := Santa{0, 0}

	revisits := 0

	seen[House{0, 0}] = true
	for i := 0; scanner.Scan(); i++ {
		char := scanner.Text()
		if char == "<" {
			if i%2 == 0 {
				santa.col--
			} else {
				roboSanta.col--
			}
		} else if char == ">" {
			if i%2 == 0 {
				santa.col++
			} else {
				roboSanta.col++
			}
		}
		if char == "^" {
			if i%2 == 0 {
				santa.row--
			} else {
				roboSanta.row--
			}
		}
		if char == "v" {
			if i%2 == 0 {
				santa.row++
			} else {
				roboSanta.row++
			}
		}

		row := santa.row
		col := santa.col
		if i%2 != 0 {
			row = roboSanta.row
			col = roboSanta.col
		}
		house := House{row, col}

		if _, ok := seen[house]; !ok {
			seen[house] = true
		} else {
			revisits++
		}

	}
	fmt.Println(len(seen))
	fmt.Println(revisits)
}
