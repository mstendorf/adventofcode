package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var password = []byte("abcdefgh")

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "swap":
			if parts[1] == "position" {
				swapPosition(parts[2], parts[5])
			} else {
				swapLetter(parts[2], parts[5])
			}
		case "rotate":
			if parts[1] == "based" {
				rotateBased(parts[6])
			} else {
				rotate(parts[1], parts[2])
			}
		case "reverse":
			reverse(parts[2], parts[4])
		case "move":
			move(parts[2], parts[5])
		default:
			panic("Unknown command")
		}

	}
	fmt.Println("Part 1", string(password))
}

func swapPosition(x, y string) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	password[ix], password[iy] = password[iy], password[ix]
}

func swapLetter(x, y string) {
	ix, iy := 0, 0
	for i, c := range password {
		if c == x[0] {
			ix = i
		}
		if c == y[0] {
			iy = i
		}
	}
	password[ix], password[iy] = password[iy], password[ix]
}

func rotate(dir, steps string) {
	stepsInt := int(steps[0] - '0')
	if dir == "left" {
		stepsInt = len(password) - stepsInt
	}
	newPassword := make([]byte, len(password))
	for i, c := range password {
		newPassword[(i+stepsInt)%len(password)] = c
	}
	copy(password, newPassword)
}

func rotateBased(x string) {
	ix := 0
	for i, c := range password {
		if c == x[0] {
			ix = i
			break
		}
	}
	rotations := 1 + ix
	if ix >= 4 {
		rotations++
	}
	rotate("right", strconv.Itoa(rotations))
}

func reverse(x, y string) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	for ix < iy {
		password[ix], password[iy] = password[iy], password[ix]
		ix++
		iy--
	}
}

func move(x, y string) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	if ix < iy {
		for ix < iy {
			password[ix], password[ix+1] = password[ix+1], password[ix]
			ix++
		}
	} else {
		for ix > iy {
			password[ix], password[ix-1] = password[ix-1], password[ix]
			ix--
		}
	}
}
