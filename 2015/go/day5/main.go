package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(part1(file))
	file.Seek(0, 0)
	fmt.Println(part2(file))

}

func part1(file *os.File) int {

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			count++
		}

	}
	return count

}

func part2(file *os.File) int {

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isNice2(line) {
			count++
		}

	}
	return count
}

func isNice2(line string) bool {

	if hasPairTwice(line) && hasRepeatingLetter(line) {
		return true
	}

	return false
}

func hasPairTwice(line string) bool {

	for i := 0; i < len(line)-1; i++ {
		if strings.Contains(line[i+2:], line[i:i+2]) {
			return true
		}
	}

	return false
}

func hasRepeatingLetter(line string) bool {

	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}

	return false
}

func isNice(line string) bool {

	if hasThreeVowels(line) && hasDoubleLetter(line) && !hasForbiddenStrings(line) {
		return true
	}

	return false
}

func hasThreeVowels(line string) bool {

	vowels := "aeiou"
	count := 0
	for _, c := range line {
		if strings.Contains(vowels, string(c)) {
			count++
			if count == 3 {
				return true
			}
		}
	}

	return false
}

func hasDoubleLetter(line string) bool {

	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}

	return false
}

func hasForbiddenStrings(line string) bool {

	for _, s := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(line, s) {
			return true
		}
	}

	return false
}
