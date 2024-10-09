package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	part1(file)
	part2(file)
}

func part1(file *os.File) {

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	code := 0
	memory := 0

	for scanner.Scan() {
		line := scanner.Text()
		c, m := count(line)
		code += c
		memory += m
	}
	fmt.Println("part1")
	fmt.Println(code-memory, code, memory)
}

func part2(file *os.File) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	code := 0
	expanded := 0

	// expandedCount(`""`)
	// expandedCount(`"abc"`)
	// expandedCount(`"aaa\"aaa"`)
	// expandedCount(`"\x27"`)
	for scanner.Scan() {
		line := scanner.Text()
		code += len(line)
		expanded += expandedCount(line)
	}
	fmt.Println("part2")
	fmt.Println(expanded-code, code, expanded)
}

func expandedCount(str string) int {
	expanded := ""
	for i := 0; i < len(str); i++ {
		// fmt.Println("char", string(str[i]))
		if str[i] == '\\' {
			if str[i+1] == '\\' {
				expanded += "\\\\\\\\"
				i++
			} else if str[i+1] == '"' {
				expanded += "\\\\\\\""
				i++
			} else if str[i+1] == 'x' {
				expanded += "\\\\x"
				i++
			}
		} else if str[i] == '"' {
			expanded += "\\\""
		} else {
			expanded += string(str[i])
		}
		// fmt.Println(expanded)
	}

	// println(str, expanded, len(str), len(expanded)+2)
	return len(expanded) + 2
}

func count(s string) (int, int) {
	str := s[1 : len(s)-1]

	code := len(s)
	memory := 0

	for i := 0; i < len(str); i++ {
		memory++
		if str[i] == '\\' {
			if str[i+1] == '\\' {
				i++
			} else if str[i+1] == '"' {
				i++
			} else if str[i+1] == 'x' {
				i += 3
			}
		}
	}
	return code, memory
}
