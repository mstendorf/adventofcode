package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	clues := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}

	part1(file, clues)
	part2(file, clues)
}

func part1(file *os.File, clues map[string]int) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		sue := matches[1]
		clue1 := matches[2]
		clue1Value, _ := strconv.Atoi(matches[3])
		clue2 := matches[4]
		clue2Value, _ := strconv.Atoi(matches[5])
		clue3 := matches[6]
		clue3Value, _ := strconv.Atoi(matches[7])

		if clues[clue1] == clue1Value {
			if clues[clue2] == clue2Value {
				if clues[clue3] == clue3Value {
					fmt.Println(sue)
				}
			}
		}
	}
}

func part2(file *os.File, clues map[string]int) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
outer:
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		sue := matches[1]
		sueClues := map[string]int{}
		sueClues[matches[2]], _ = strconv.Atoi(matches[3])
		sueClues[matches[4]], _ = strconv.Atoi(matches[5])
		sueClues[matches[6]], _ = strconv.Atoi(matches[7])

		for clue, value := range sueClues {
			if clue == "cats" || clue == "trees" {
				if value <= clues[clue] {
					continue outer
				}
			} else if clue == "pomeranians" || clue == "goldfish" {
				if value >= clues[clue] {
					continue outer
				}
			} else {
				if value != clues[clue] {
					continue outer
				}
			}
		}
		fmt.Println(sue)
	}

}
