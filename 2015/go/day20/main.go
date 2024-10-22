package main

type House struct {
	num_presents int
}

func main() {

	num_packages := 34000000

	println(part1(num_packages))
	println(part2(num_packages))

}

func part1(num_packages int) int {
	max_iterations := num_packages / 10
	houses := make([]House, max_iterations)
	for i := 1; i < max_iterations; i++ {
		for j := i; j < max_iterations; j += i {
			houses[j].num_presents += i * 10
		}
	}

	for i := 1; i < max_iterations; i++ {
		if houses[i].num_presents >= num_packages {
			return i
		}
	}
	return 0
}

func part2(num_packages int) int {
	max_iterations := num_packages / 10
	houses := make([]House, max_iterations)
	for i := 1; i < max_iterations; i++ {
		for j := i; j < max_iterations && j <= i*50; j += i {
			houses[j].num_presents += i * 11
		}
	}

	for i := 1; i < max_iterations; i++ {
		if houses[i].num_presents >= num_packages {
			return i
		}
	}
	return 0
}
