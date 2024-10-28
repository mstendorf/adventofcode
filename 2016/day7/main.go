package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re_in_bracket := regexp.MustCompile(`\[\w+\]`)

	count := 0
	abacount := 0
	for scanner.Scan() {
		line := scanner.Text()

		in_bracket := re_in_bracket.FindAllString(line, -1)
		line_trimmed := re_in_bracket.ReplaceAllString(line, " ")

		abba := hasABBA(line_trimmed)
		if abba {
			skip := false
			for _, s := range in_bracket {
				a := hasABBA(s)
				if a {
					skip = true
					break
				}
			}
			if !skip {
				count++
			}
		}

		if hasABAset(line_trimmed, in_bracket, line) {
			abacount++
		}
	}

	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", abacount)
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func hasABAset(s string, brackets []string, full string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			aba := s[i : i+3]
			bab := fmt.Sprintf("%c%c%c", aba[1], aba[0], aba[1])
			for _, bracket := range brackets {
				if strings.Contains(bracket, bab) {
					return true
				}
			}
		}
	}
	return false
}
