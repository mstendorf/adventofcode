package main

import (
	"fmt"
)

type Elf struct {
	id          int
	has_present bool
	next        *Elf
}

func main() {
	num_elves := 3014387
	// num_elves := 5
	fmt.Println("Part 1", part1(num_elves))
	fmt.Println("Part 2", part2(num_elves))
	part2Fancy(num_elves)

}

func part1(num_elves int) int {
	head := &Elf{1, true, nil}
	current := head
	for i := 2; i <= num_elves; i++ {
		current.next = &Elf{i, true, nil}
		current = current.next
	}

	current.next = head
	current = head
	// Part 1
	for current.next != nil {

		if current.next.has_present {
			current.next.has_present = false
		} else {
			current.next = current.next.next
		}
		if current.next == current {
			break
		}
		current = current.next
	}
	return current.id
}

func part2Fancy(num int) {
	i := 1

	for i*3 < num {
		i *= 3
	}
	fmt.Println(num - i)
}

func part2(num_elves int) int {
	// this idea with creating two dequeues and rotating them is stolen
	left := Dequeue{}
	right := Dequeue{}
	for i := 1; i <= num_elves/2; i++ {
		left.Push(i)
	}
	for i := num_elves/2 + 1; i <= num_elves; i++ {
		right.Pushleft(i)
	}

	for left.length > 1 || right.length > 1 {
		if left.length > right.length {
			left.Pop()
		} else {
			right.Pop()
		}

		// rotate the elves to simulate the circle
		right.Pushleft(left.Popleft())
		left.Push(right.Pop())
	}
	return left.Pop()
}
