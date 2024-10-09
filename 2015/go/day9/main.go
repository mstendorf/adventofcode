package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	part1(file)
	part2(file)
}

func buildCityMap(file *os.File) map[string]map[string]int {
	file.Seek(0, 0)
	cities := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		citys := strings.Split(parts[0], " to ")
		city1 := citys[0]
		city2 := citys[1]
		distance, _ := strconv.Atoi(parts[1])
		if cities[city1] == nil {
			cities[city1] = make(map[string]int)
		}
		if cities[city2] == nil {
			cities[city2] = make(map[string]int)
		}
		cities[city1][city2] = distance
		cities[city2][city1] = distance

	}
	return cities
}

func part1(file *os.File) {
	cities := buildCityMap(file)
	distance := travelShort(cities, "Faerun", []string{"Faerun"})
	fmt.Println(distance)
}

func part2(file *os.File) {
	// should probs have anticipated this and just used permutations for part 1 instead
	// of writing the greedy solution
	cities := buildCityMap(file)
	cityList := make([]string, 0)
	for city := range cities {
		cityList = append(cityList, city)
	}
	perms := permutations(cityList)
	maxDistance := 0
	for _, perm := range perms {
		distance := 0
		for i := 0; i < len(perm)-1; i++ {
			distance += cities[perm[i]][perm[i+1]]
		}
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	fmt.Println(maxDistance)
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

func travelShort(cities map[string]map[string]int, currentCity string, visitedCities []string) int {
	// find the shortest path to visit all cities
	if len(visitedCities) == len(cities) {
		return 0
	}
	minDistance := 1000000
	for city, distance := range cities[currentCity] {
		if !slices.Contains(visitedCities, city) {
			newVisitedCities := append(visitedCities, city)
			newDistance := travelShort(cities, city, newVisitedCities)
			if newDistance+distance < minDistance {
				minDistance = newDistance + distance
			}
		}
	}
	return minDistance
}
