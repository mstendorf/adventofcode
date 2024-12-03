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

	var password = []byte("abcdefgh")
	fmt.Println("Part 1", string(part1(file, password)))
	password = []byte("fbgdceah")
	part2(file, password)
}

func part1(file *os.File, password []byte) []byte {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		handleCommand(parts, &password)

	}
	return password
}

// the one time i decide to be lazy and not parse the file to some data structure it bites me in the ass
// got even more lazy and just kept reading the file. Not efficient or smart, but it works
func part2(file *os.File, password []byte) {
	file.Seek(0, 0)
	perms := permutations(password)
	for _, line := range perms {
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		if string(part1(file, lineCopy)) == string(password) {
			fmt.Println("Part 2", string(line))
			break
		}
	}

}

func permutations(arr []byte) [][]byte {
	if len(arr) == 1 {
		return [][]byte{arr}
	}
	perms := make([][]byte, 0)
	for i, val := range arr {
		rest := make([]byte, len(arr)-1)
		copy(rest, arr[:i])
		copy(rest[i:], arr[i+1:])
		for _, p := range permutations(rest) {
			perm := make([]byte, len(arr))
			perm[0] = val
			copy(perm[1:], p)
			perms = append(perms, perm)
		}
	}
	return perms
}

func handleCommand(parts []string, password *[]byte) {
	switch parts[0] {
	case "swap":
		if parts[1] == "position" {
			swapPosition(parts[2], parts[5], password)
		} else {
			swapLetter(parts[2], parts[5], password)
		}
	case "rotate":
		if parts[1] == "based" {
			rotateBased(parts[6], password)
		} else {
			rotate(parts[1], parts[2], password)
		}
	case "reverse":
		reverse(parts[2], parts[4], password)
	case "move":
		move(parts[2], parts[5], password)
	default:
		panic("Unknown command")
	}
}

func swapPosition(x, y string, password *[]byte) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	(*password)[ix], (*password)[iy] = (*password)[iy], (*password)[ix]
}

func swapLetter(x, y string, password *[]byte) {
	ix, iy := 0, 0
	for i, c := range *password {
		if c == x[0] {
			ix = i
		}
		if c == y[0] {
			iy = i
		}
	}
	(*password)[ix], (*password)[iy] = (*password)[iy], (*password)[ix]
}

func rotate(dir, steps string, password *[]byte) {
	stepsInt := int(steps[0] - '0')
	if dir == "left" {
		stepsInt = len((*password)) - stepsInt
	}
	newPassword := make([]byte, len((*password)))
	for i, c := range *password {
		newPassword[(i+stepsInt)%len((*password))] = c
	}
	copy((*password), newPassword)
}

func rotateBased(x string, password *[]byte) {
	ix := 0
	for i, c := range *password {
		if c == x[0] {
			ix = i
			break
		}
	}
	rotations := 1 + ix
	if ix >= 4 {
		rotations++
	}
	rotate("right", strconv.Itoa(rotations), password)
}

func reverse(x, y string, password *[]byte) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	for ix < iy {
		(*password)[ix], (*password)[iy] = (*password)[iy], (*password)[ix]
		ix++
		iy--
	}
}

func move(x, y string, password *[]byte) {
	ix, iy := int(x[0]-'0'), int(y[0]-'0')
	if ix < iy {
		for ix < iy {
			(*password)[ix], (*password)[ix+1] = (*password)[ix+1], (*password)[ix]
			ix++
		}
	} else {
		for ix > iy {
			(*password)[ix], (*password)[ix-1] = (*password)[ix-1], (*password)[ix]
			ix--
		}
	}
}
