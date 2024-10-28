package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Part 1:", decompress(text, false))
		fmt.Println("Part 2:", decompress(text, true))
	}
}

func decompress(s string, recursive bool) int {
	if !strings.Contains(s, "(") {
		return len(s)
	}

	count := 0
	for strings.Contains(s, "(") {
		next_idx := strings.Index(s, "(")
		count += next_idx

		s = s[next_idx:]

		marker_end := strings.Index(s, ")")
		marker := s[1:marker_end]

		s = s[marker_end+1:]

		parts := strings.Split(marker, "x")
		chars, _ := strconv.Atoi(parts[0])
		amount, _ := strconv.Atoi(parts[1])

		if recursive {
			count += decompress(s[:chars], true) * amount
		} else {

			count += chars * amount
		}

		s = s[chars:]
	}
	count += len(s)
	return count
}
