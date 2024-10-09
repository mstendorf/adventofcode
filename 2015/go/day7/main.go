package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Operator uint8

const (
	MOV Operator = iota
	AND
	OR
	LSHIFT
	RSHIFT
	NOT
)

type Gate struct {
	op       Operator
	in1, in2 string
}

var (
	gates     = make(map[string]Gate)
	connected = make(map[string]uint16)
)

func backtrack(dst string) uint16 {
	// part 2 override, comment this if for part 1
	if dst == "b" {
		return 16076
	}
	if v, err := strconv.ParseUint(dst, 10, 16); err == nil {
		return uint16(v)
	}
	if v, ok := connected[dst]; ok {
		return v
	}

	gate := gates[dst]
	var v uint16

	switch gate.op {
	case MOV:
		v = backtrack(gate.in1)
	case AND:
		v = backtrack(gate.in1) & backtrack(gate.in2)
	case OR:
		v = backtrack(gate.in1) | backtrack(gate.in2)
	case LSHIFT:
		v = backtrack(gate.in1) << backtrack(gate.in2)
	case RSHIFT:
		v = backtrack(gate.in1) >> backtrack(gate.in2)
	case NOT:
		v = ^backtrack(gate.in1)
	}

	connected[dst] = v
	return v
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var (
			op            Operator
			in1, in2, dst string
		)

		if n, _ := fmt.Sscanf(line, "%s -> %s", &in1, &dst); n == 2 {
			op = MOV
		} else if n, _ := fmt.Sscanf(line, "%s AND %s -> %s", &in1, &in2, &dst); n == 3 {
			op = AND
		} else if n, _ := fmt.Sscanf(line, "%s OR %s -> %s", &in1, &in2, &dst); n == 3 {
			op = OR
		} else if n, _ := fmt.Sscanf(line, "%s LSHIFT %s -> %s", &in1, &in2, &dst); n == 3 {
			op = LSHIFT
		} else if n, _ := fmt.Sscanf(line, "%s RSHIFT %s -> %s", &in1, &in2, &dst); n == 3 {
			op = RSHIFT
		} else if n, _ := fmt.Sscanf(line, "NOT %s -> %s", &in1, &dst); n == 2 {
			op = NOT
		} else {
			panic(line)
		}

		gates[dst] = Gate{op, in1, in2}
	}

	fmt.Println(backtrack("a"))

}
