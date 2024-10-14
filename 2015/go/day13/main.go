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
		os.Exit(1)
	}
	defer file.Close()

	score := part1(file)
	fmt.Println(score)
	score2 := part2(file)
	fmt.Println(score2)
}

func part1(file *os.File) int {
	personMap := buildPersonMap(file)
	persons := personList(personMap)
	personPerms := permutations(persons)

	return findHighestHappiness(personMap, personPerms)
}

func part2(file *os.File) int {
	personMap := buildPersonMapWithMe(file)
	persons := personList(personMap)
	fmt.Println(persons)
	// persons = append(persons, "me")
	personPerms := permutations(persons)

	return findHighestHappiness(personMap, personPerms)
}

func findHighestHappiness(personMap map[string]map[string]int, personPerms [][]string) int {
	highest := 0
	for _, perm := range personPerms {
		happiness := 0
		for i := 0; i < len(perm); i++ {
			person1 := perm[i]
			person2 := perm[(i+1)%len(perm)]
			happiness += personMap[person1][person2]
			happiness += personMap[person2][person1]
		}
		if happiness > highest {
			highest = happiness
		}
	}
	return highest
}

func permutations(arr []string) [][]string {
	if len(arr) == 1 {
		return [][]string{arr}
	}
	perms := make([][]string, 0)
	for i, val := range arr {
		rest := make([]string, len(arr)-1)
		copy(rest, arr[:i])
		copy(rest[i:], arr[i+1:])
		for _, p := range permutations(rest) {
			perm := make([]string, len(arr))
			perm[0] = val
			copy(perm[1:], p)
			perms = append(perms, perm)
		}
	}
	return perms
}
func personList(personMap map[string]map[string]int) []string {
	persons := make([]string, len(personMap))
	i := 0
	for k := range personMap {
		persons[i] = k
		i++
	}
	return persons
}

func buildPersonMapWithMe(file *os.File) map[string]map[string]int {
	file.Seek(0, 0)
	persons := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		person1 := matches[1]
		effect := matches[2]
		happiness, _ := strconv.Atoi(matches[3])
		person2 := matches[4]
		if persons[person1] == nil {
			persons[person1] = make(map[string]int)
		}
		// if persons[person2] == nil {
		// 	persons[person2] = make(map[string]int)
		// }
		if effect == "gain" {
			persons[person1][person2] = happiness
		} else {
			persons[person1][person2] = -happiness
		}

		if persons["me"] == nil {
			persons["me"] = make(map[string]int)
		}

		persons["me"][person1] = 0
		persons[person1]["me"] = 0

	}
	return persons
}

func buildPersonMap(file *os.File) map[string]map[string]int {
	file.Seek(0, 0)
	persons := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		person1 := matches[1]
		effect := matches[2]
		happiness, _ := strconv.Atoi(matches[3])
		person2 := matches[4]
		if persons[person1] == nil {
			persons[person1] = make(map[string]int)
		}
		// if persons[person2] == nil {
		// 	persons[person2] = make(map[string]int)
		// }
		if effect == "gain" {
			persons[person1][person2] = happiness
		} else {
			persons[person1][person2] = -happiness
		}

	}
	return persons
}
