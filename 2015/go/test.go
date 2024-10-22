package main

import "fmt"

func main() {

	part1()
}

func part1() {
	arr := makeSlice()
	for i := 0; i < 100; i++ {
		arr = useSlice(arr)
	}

	fmt.Println(arr)
}

func makeSlice() []int {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return arr
}

func useSlice(arr []int) []int {
	arr3 := make([]int, 10)
	copy(arr3, arr)
	arr3[0] = arr[0] + 1
	arr[0] = 11
	fmt.Println(arr3)
	return arr3
}
