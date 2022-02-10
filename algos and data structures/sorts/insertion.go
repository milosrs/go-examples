package main

import "fmt"

func main() {
	x := []int{50, 12, 123457, 8537, 7235, 1232, 7532, 8755, 803}
	fmt.Printf("Unsorted: %v\n", x)
	fmt.Printf("Sorted INC: %v\n", increasingInsertionSort(x))
	fmt.Printf("Sorted DEC: %v\n", decreasingInsertionSort(x))
}

func increasingInsertionSort(arr []int) []int {
	i := 0

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i = j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}

	return arr
}

func decreasingInsertionSort(arr []int) []int {
	i := 0

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i = j - 1
		for i >= 0 && arr[i] < key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}

	return arr
}
