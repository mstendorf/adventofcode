package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := part1(string(file))
	fmt.Println(sum)
	sum2 := part2(&file)
	fmt.Println(sum2)

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
func part2(contents *[]byte) float64 {
	// this is inspired if not stolen from online posts, learned some nifty tricks so decided to keep it
	var f interface{}
	var output float64
	json.Unmarshal(*contents, &f)

	output = parse(f)

	return output
}

func parse(f interface{}) (output float64) {
outer:
	switch fv := f.(type) {
	case []interface{}:
		for _, val := range fv {
			output += parse(val)
		}
	case float64:
		output += fv
	case map[string]interface{}:
		for _, val := range fv {
			if val == "red" {
				break outer
			}
		}
		for _, val := range fv {
			output += parse(val)
		}
	}

	return output
}
