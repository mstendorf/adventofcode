package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Replacement struct {
	From string
	To   string
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`(\w+) => (\w+)`)

	replacements := make([]Replacement, 0)
	subject := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			subject = line
			continue
		}

		rep := Replacement{matches[1], matches[2]}
		replacements = append(replacements, rep)
	}
	fmt.Println(subject)
	part1(replacements, subject)
	part2(replacements, subject)
}

func part1(replacements []Replacement, subject string) {
	// Find all possible replacements
	possibilities := make(map[string]bool)
	for _, rep := range replacements {
		re := regexp.MustCompile(rep.From)
		matches := re.FindAllStringIndex(subject, -1)
		for _, match := range matches {
			newSubject := subject[:match[0]] + rep.To + subject[match[1]:]
			possibilities[newSubject] = true
		}
	}
	fmt.Println(len(possibilities))
}

func part2(replacements []Replacement, subject string) {
	// I feel like this might be luck, but i am shit out of luck to write something clever that guarantees
	// a correct solution for any given input
	n := 0
	for subject != "e" {
		for _, rep := range replacements {
			re := regexp.MustCompile(rep.To)
			matches := re.FindAllStringIndex(subject, 1)
			for _, match := range matches {
				subject = subject[:match[0]] + rep.From + subject[match[1]:]
				n++
			}
		}
	}
	fmt.Println(n)
}
