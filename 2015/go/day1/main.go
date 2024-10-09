package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	s := bufio.NewScanner(strings.NewReader(string(data)))
	s.Split(bufio.ScanRunes)

	floor := 0
	for i := 1; s.Scan(); i++ {
		if s.Text() == "(" {
			floor++
		} else if s.Text() == ")" {
			floor--
		}
		if floor == -1 {
			fmt.Println("Basement at:", i)
		}
	}
	fmt.Println("Floor:", floor)

}
