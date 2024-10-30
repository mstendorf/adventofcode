package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Disc struct {
	positions int
	current   int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Disc #\d has (\d+)\s\w+; at time=0, it is at position (\d+).`)
	discs := make([]Disc, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		position, _ := strconv.Atoi(matches[1])
		current, _ := strconv.Atoi(matches[2])

		discs = append(discs, Disc{positions: position, current: current})
	}

	fmt.Println("Part 1", runClock(discs))
	discs = append(discs, Disc{positions: 11, current: 0})
	fmt.Println("Part 2", runClock(discs))
}

func runClock(discs []Disc) int {
	time := 0
	for {
		if isAligned(discs, time) {
			break
		}
		time++
	}
	return time
}

func isAligned(discs []Disc, time int) bool {
	for i, disc := range discs {
		if (disc.current+time+i+1)%disc.positions != 0 {
			return false
		}
	}
	return true
}
