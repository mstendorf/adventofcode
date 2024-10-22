package main

import (
	"bufio"
	"log"
	"math"
	"math/bits"
	"os"
	"slices"
	"strconv"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	packages := []int{}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		sum += num
		packages = append(packages, num)
	}
	// balance(packages, 3)
	balance(packages, 4)

}

func balance(containers []int, compartments int) {

	compartment_size := sum(containers) / compartments

	min_len := len(containers)
	min_qe := math.MaxInt64
	for i := 1; i < len(containers); i++ {

		if min_len < i {
			break
		}
		combos := Combinations(containers, i)
		for _, combo := range combos {
			if len(combo) > min_len {
				continue
			}
			if sum(combo) == compartment_size {
				cp := make([]int, len(containers))
				copy(cp, containers)
				rest := remove(cp, combo)
				if canSplit(rest, compartment_size, compartments-1) {
					if len(combo) <= min_len {
						min_len = len(combo)
						qe := product(combo)
						if qe < min_qe {
							min_qe = qe
							// fmt.Println(min_len, min_qe)
						}

					}
				}

			}
		}
	}
	println(min_qe)
}

func canSplit(containers []int, compartment_size, compartments int) bool {

	// fmt.Println("--------------------------")
	cp := make([]int, len(containers))
	copy(cp, containers)
	// fmt.Println("can split", cp, compartment_size, compartments)
	for i := 1; i < len(cp); i++ {
		combos := Combinations(cp, i)
		for _, combo := range combos {
			s := sum(combo)
			if s == compartment_size {
				// fmt.Println(i)
				rest := remove(cp, combo)

				if compartments == 2 && sum(rest) == compartment_size {
					// fmt.Println("2 found ", combo, rest)
					return true
				} else if canSplit(rest, compartment_size, compartments-1) {
					// fmt.Println("can split", combo)
					return true
				}
			}
		}
	}
	// fmt.Println("--------------------------")
	return false
}

func remove(slice []int, slice2 []int) []int {
	// fmt.Println("-------------------------- remove")
	// fmt.Println(slice, slice2)
	rest := slices.DeleteFunc(slice, func(x int) bool {
		for _, y := range slice2 {
			if x == y {
				return true
			}
		}
		return false
	})
	// fmt.Println(rest)
	// fmt.Println("-------------------------- remove end")
	return rest
}

func product(containers []int) int {
	product := 1
	for _, container := range containers {
		product *= container
	}
	return product
}

func sum(containers []int) int {
	sum := 0
	for _, container := range containers {
		sum += container
	}
	return sum
}

func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
