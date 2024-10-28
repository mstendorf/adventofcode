package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Bot struct {
	id      int
	values  []int
	low     int
	lowBot  bool
	high    int
	highBot bool
}

func (b *Bot) addValue(value int) {
	b.values = append(b.values, value)
}

func (b *Bot) hasTwoValues() bool {
	return len(b.values) == 2
}

func (b *Bot) giveValues() (int, int) {
	return slices.Min(b.values), slices.Max(b.values)
}

type Instruction struct {
	botID int
	value int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bots := make(map[int]*Bot)
	instructions := []Instruction{}
	outputs := make(map[int]int)
	scanner := bufio.NewScanner(file)
	value := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	bot := regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

	for scanner.Scan() {

		valueMatches := value.FindStringSubmatch(scanner.Text())
		if len(valueMatches) > 0 {
			// fmt.Println(valueMatches)
			botID, _ := strconv.Atoi(valueMatches[2])
			value, _ := strconv.Atoi(valueMatches[1])
			inst := Instruction{botID: botID, value: value}
			instructions = append(instructions, inst)
		} else {

			botMatches := bot.FindStringSubmatch(scanner.Text())
			botID, _ := strconv.Atoi(botMatches[1])
			lowID, _ := strconv.Atoi(botMatches[3])
			highID, _ := strconv.Atoi(botMatches[5])
			lowBot := botMatches[2] == "bot"
			highBot := botMatches[4] == "bot"

			if _, ok := bots[botID]; !ok {
				bots[botID] = &Bot{id: botID}
			}
			bots[botID].low = lowID
			bots[botID].high = highID
			bots[botID].lowBot = lowBot
			bots[botID].highBot = highBot
		}

	}
	part1(bots, instructions, &outputs)
	part2(&outputs)
}

func part1(bots map[int]*Bot, instructions []Instruction, outputs *map[int]int) {
	for _, inst := range instructions {
		bots[inst.botID].addValue(inst.value)
		execute(inst.botID, bots, outputs)
	}
}

func part2(outputs *map[int]int) {
	// fmt.Println(*outputs)
	fmt.Println("Part 2:", (*outputs)[0]*(*outputs)[1]*(*outputs)[2])
}

func execute(bot int, bots map[int]*Bot, outputs *map[int]int) {

	if bots[bot].hasTwoValues() {

		low, high := bots[bot].giveValues()
		if low == 17 && high == 61 {
			fmt.Println("Part 1:", bot)
		}
		if bots[bot].lowBot {
			if _, ok := bots[bots[bot].low]; !ok {
				bots[bots[bot].low] = &Bot{id: bots[bot].low}
			}
			bots[bots[bot].low].addValue(low)
			execute(bots[bot].low, bots, outputs)
		} else {
			// fmt.Println("output", bot, low)
			(*outputs)[bots[bot].low] = low
		}
		if bots[bot].highBot {
			if _, ok := bots[bots[bot].high]; !ok {
				bots[bots[bot].high] = &Bot{id: bots[bot].high}
			}
			bots[bots[bot].high].addValue(high)
			execute(bots[bot].high, bots, outputs)
		} else {
			// fmt.Println("output", bot, high)
			(*outputs)[bots[bot].high] = high
		}

		bots[bot].values = []int{}
	}
}
