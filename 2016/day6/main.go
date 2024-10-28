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

	characters := make([][]rune, 0)
	// characters = append(characters, make([]string, 0))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(characters) < len(line) {
			for i := 0; i < len(line); i++ {
				characters = append(characters, make([]rune, 0))
			}
		}
		for i, char := range line {
			// fmt.Println(i, string(char))
			if characters[i] == nil {
				characters[i] = make([]rune, 0)
			}
			// fmt.Println("aftger")

			characters[i] = append(characters[i], rune(char))
		}
	}
	phraseMax := ""
	phraseMin := ""
	for _, char := range characters {
		phraseMax += string(mostCommonChar(char))
		phraseMin += string(leastCommonChar(char))
	}
	fmt.Println("Part 1", phraseMax)
	fmt.Println("Part 2", phraseMin)
}

func leastCommonChar(roomCode []rune) rune {
	counts := make(map[rune]int)
	for _, char := range roomCode {
		counts[char]++
	}
	min := 10000
	var char rune
	for k, v := range counts {
		if v < min {
			min = v
			char = k
		}
	}
	return char
}

func mostCommonChar(roomCode []rune) rune {
	counts := make(map[rune]int)
	for _, char := range roomCode {
		counts[char]++
	}
	max := 0
	var char rune
	for k, v := range counts {
		if v > max {
			max = v
			char = k
		}
	}
	return char
}
