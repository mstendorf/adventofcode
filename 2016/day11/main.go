package main

import (
	"fmt"
	"slices"
	"sort"
)

type State struct {
	items         []int
	graphLocation int
}

func (s State) Copy() State {
	newState := State{make([]int, len(s.items)), s.graphLocation}
	copy(newState.items, s.items)
	return newState
}

func main() {
	// this was retardedly hard for me, took way longer than i care to admit
	// the data structure is inspired by some tips i found online
	// this is very inefficient and takes a while to run, but it works.
	// there is some easy optimizations to be made, but i'm not going to bother

	// [elevator, generator1,chips1, generator2, chips2, generator3...]
	state := []int{0, 0, 0, 1, 2, 1, 2, 1, 2, 1, 2}
	fmt.Println("Part 1:", solve(state))

	state2 := []int{0, 0, 0, 1, 2, 1, 2, 1, 2, 1, 2, 0, 0, 0, 0}
	fmt.Println("Part 2:", solve(state2))

}

func solve(floors []int) int {

	visited_states := make([][]int, 0)
	// [elevator, generator1,chips1, generator2, chips2, generator3...]
	visited := []State{State{floors, 0}}
	// visited := []State{State{[]int{0, 1, 0, 2, 0}, 0}}
	moves := 0

	for i := 0; i < len(visited); i++ {

		state := visited[i]
		items_on_floor := findOnFloor(state)

		possible_moves := combinations(items_on_floor)
		directions := possibleDirections(state.items[0])

		current_try := state.graphLocation

		//copy old state to base all new moves on
		new_state := state.Copy()

		for _, direction := range directions {
			for _, move := range possible_moves {
				// fmt.Println("move", move)
				moves++

				test_state := new_state.Copy()

				// move the elevator
				test_state.items[0] += direction

				if len(move) == 2 {
					// take two items
					test_state.items[move[0]] += direction
					test_state.items[move[1]] += direction
				} else {
					// take one item
					test_state.items[move[0]] += direction
				}

				// check if state has already been visited
				// if not, add it to the list of visited states
				if isVisited(test_state, &visited_states) {
					continue
				}

				// check if state is valid
				if isValid(test_state) {
					test_state.graphLocation = current_try + 1
					visited = append(visited, test_state)
				}

				sum := 0
				for _, v := range test_state.items {
					sum += v
				}

				if sum == len(test_state.items)*3 {
					return current_try + 1

				}
			}
		}
	}
	return -1
}

func generate_chip_pair(state State) [][]int {
	items := state.items[1:]
	var pairs [][]int
	for i := 0; i < len(items)-1; i += 2 {
		pairs = append(pairs, []int{items[i], items[i+1]})
	}
	return pairs
}

func isValid(state State) bool {
	pairs := generate_chip_pair(state)
	generators := []int{}
	for _, pair := range pairs {
		generators = append(generators, pair[0])
	}

	for _, pair := range pairs {
		if pair[0] != pair[1] {
			if slices.Contains(generators, pair[1]) {
				return false
			}

		}
	}
	return true
}

func isVisited(state State, visited *[][]int) bool {
	pairs := generate_chip_pair(state)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	// unpack pairs to unique slice
	var unique []int
	unique = append(unique, state.items[0])
	for _, pair := range pairs {

		unique = append(unique, pair...)
	}

	for _, v := range *visited {
		if slices.Equal(unique, v) {
			return true
		}
	}
	*visited = append(*visited, unique)
	return false
}

func possibleDirections(floor int) []int {
	if floor == 0 {
		return []int{1}
	}
	if floor == 3 {
		return []int{-1}
	}
	return []int{1, -1}
}

func findOnFloor(state State) []int {
	var items []int
	for i := 0; i < len(state.items); i++ {
		if state.items[i] == state.items[0] {
			items = append(items, i)
		}
	}
	return items[1:]
}

func combinations(items []int) [][]int {
	var combinations [][]int
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			combinations = append(combinations, []int{items[i], items[j]})
		}
	}
	// append single items
	for _, item := range items {
		combinations = append(combinations, []int{item})
	}
	return combinations
}
