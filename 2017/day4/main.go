package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validCount := 0
	validCountPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if isValid(words) {
			validCount++
		}
		if isValidPart2(words) {
			validCountPart2++
		}
	}
	fmt.Println("Part 1", validCount)
	fmt.Println("Part 2", validCountPart2)
}

func isValid(words []string) bool {
	seen := make(map[string]bool)
	for _, word := range words {
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}

func isValidPart2(words []string) bool {
	seen := make(map[string]bool)
	for _, word := range words {
		sorted := sortString(word)
		if seen[sorted] {
			return false
		}
		seen[sorted] = true
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
