package main

import "fmt"

func main() {
	x := []int{50, 12, 123457, 8537, 7235, 1232, 7532, 8755, 803}
	fmt.Printf("Unsorted: %v\n", x)
	fmt.Printf("Sorted INC: %v\n", selectionSortInc(x))
	fmt.Printf("Sorted DEC: %v\n", selectionSortDec(x))
}

func min(a []int, start int) (int, int) {
	for j := start; j < len(a); j++ {
		if a[j] < a[start] {
			start = j
		}
	}

	return a[start], start
}

func max(a []int, start int) (int, int) {
	for j := start; j < len(a); j++ {
		if a[j] > a[start] {
			start = j
		}
	}

	return a[start], start
}

func selectionSortInc(a []int) []int {
	end := len(a)

	for i := 0; i < end; i++ {
		_, minInd := min(a, i)
		a[i], a[minInd] = a[minInd], a[i]
	}

	return a
}

func selectionSortDec(a []int) []int {
	end := len(a)

	for i := 0; i < end; i++ {
		_, maxInd := max(a, i)
		a[i], a[maxInd] = a[maxInd], a[i]
	}

	return a
}
