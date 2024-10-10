package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1321131112"
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)

		fmt.Println(i, len(input))
	}
	fmt.Println(len(input))
}

func lookAndSay(input string) string {
	var result strings.Builder
	count := 1
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 || input[i] != input[i+1] {
			result.WriteString(fmt.Sprintf("%d%c", count, input[i]))
			count = 1
		} else {
			count++
		}
	}
	return result.String()
}
