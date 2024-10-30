package main

import "fmt"

func main() {

	input := "00111101111101000"

	part1(input)
	part2(input)
}

func part1(input string) {

	arr := fill_disk(input, 272)
	checksum := odd_checksum(arr)
	fmt.Println("Part 1", string(checksum))
}

func part2(input string) {

	arr := fill_disk(input, 35651584)
	checksum := odd_checksum(arr)
	fmt.Println("Part 2", string(checksum))
}

func fill_disk(input string, disk_size int) []byte {
	arr := []byte(input)
	// fmt.Println(string(dragon_curve(arr)))
	for len(arr) < disk_size {
		arr = dragon_curve(arr)
	}
	arr = arr[:disk_size]
	return arr
}

func odd_checksum(input []byte) []byte {
	checksum := calc_checksum(input)
	for len(checksum)%2 == 0 {
		checksum = calc_checksum(checksum)
	}
	return checksum
}

func calc_checksum(input []byte) []byte {
	output := make([]byte, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		if input[i] == input[i+1] {
			output[i/2] = '1'
		} else {
			output[i/2] = '0'
		}
	}
	return output
}

func dragon_curve(input []byte) []byte {
	output := make([]byte, len(input)*2+1)
	copy(output, input)
	output[len(input)] = '0'
	for i := 0; i < len(input); i++ {
		if input[i] == '0' {
			output[len(input)*2-i] = '1'
		} else {
			output[len(input)*2-i] = '0'
		}
		// output[len(input)*2-i] = input[i]
	}
	return output
}

func flip(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		if input[i] == '0' {
			output[i] = '1'
		} else {
			output[i] = '0'
		}
	}
	return output
}

func reverse(s []byte) []byte {
	runes := s
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}
