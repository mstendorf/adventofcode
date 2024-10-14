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
	deers := buildDeerList(file)
	part1And2(deers)
	// part2(file)
}

type Deer struct {
	name       string
	speed      int
	flyTime    int
	restTime   int
	state      string
	activeTime int
	distance   int
	score      int
}

// this is ugly but it works
func part1And2(deerList []Deer) {
	deers := deerList
	for i := 0; i < 2503; i++ {
		for j := 0; j < len(deers); j++ {
			deer := deers[j]
			if deer.state == "flying" {
				deer.activeTime--
				deer.distance += deer.speed
				if deer.activeTime == 0 {
					deer.state = "resting"
					deer.activeTime = deer.restTime
				}
			} else {
				deer.activeTime--
				if deer.activeTime == 0 {
					deer.state = "flying"
					deer.activeTime = deer.flyTime
				}
			}
			deers[j] = deer
		}

		maxDistance := 0
		leadingDeers := make([]*Deer, 0)
		for i := 0; i < len(deers); i++ {
			if deers[i].distance > maxDistance {
				maxDistance = deers[i].distance
				leadingDeers = []*Deer{&deers[i]}

			} else if deers[i].distance == maxDistance {
				leadingDeers = append(leadingDeers, &deers[i])
			}
		}
		for i := 0; i < len(leadingDeers); i++ {

			leadingDeers[i].score = leadingDeers[i].score + 1
		}
	}
	maxDistance := 0
	for _, deer := range deers {
		if deer.distance > maxDistance {
			maxDistance = deer.distance
		}
	}
	maxScore := 0
	for _, deer := range deers {
		if deer.score > maxScore {
			maxScore = deer.score
		}
	}

	fmt.Println(maxDistance)
	fmt.Println(maxScore)
}

func buildDeerList(file *os.File) []Deer {

	scanner := bufio.NewScanner(file)
	deers := make([]Deer, 0)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
		matches := re.FindStringSubmatch(line)
		deer := matches[1]
		speed, _ := strconv.Atoi(matches[2])
		flyTime, _ := strconv.Atoi(matches[3])
		restTime, _ := strconv.Atoi(matches[4])
		deers = append(deers, Deer{name: deer, speed: speed, flyTime: flyTime, restTime: restTime, state: "flying", activeTime: flyTime, distance: 0})
	}
	return deers
}
