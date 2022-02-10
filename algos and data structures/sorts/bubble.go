package main

import "fmt"

func main() {
	a := []int{3, 41, 52, 26, 38, 57, 9, 49}
	fmt.Printf("Unsorted: %v\n", a)
	fmt.Printf("Sorted: %v\n", bubbleSort(a))
}

func bubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

	return a
}
