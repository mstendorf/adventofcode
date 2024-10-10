package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := part1(string(file))
	fmt.Println(sum)
}

func part1(s string) int {
	re := regexp.MustCompile(`[-]?\d[\d]*`)

	numbers := re.FindAllString(s, -1)
	sum := 0

	for i := range numbers {
		cur, err := strconv.Atoi(numbers[i])
		sum += cur
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return sum
}
