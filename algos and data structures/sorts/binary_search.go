package main

import "fmt"

func main() {
	x := []int{50, 12, 123457, 8537, 7235, 1232, 7532, 8755, 803}
	s := 7235 //Element from array

	if doesExistBinSearch(x, s) {
		fmt.Printf("Element %d exists\n", s)
		fmt.Printf("Element %d found on index %d", s, findIndexBinSearch(x, s))
	} else {
		fmt.Println("Element doesn't exist")
	}
}

func insertionSort(arr []int) []int {
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

func findIndexBinSearch(a []int, term int) (ind int) {
	a = insertionSort(a)
	half := len(a) / 2

	switch {
	case len(a) == 0:
		ind = -1
	case a[half] > term:
		ind = findIndexBinSearch(a[:half], term)
	case a[half] > term:
		ind = findIndexBinSearch(a[half+1:], term)
		if ind >= 0 {
			ind += half + 1
		}
	case a[half] == term:
		return ind
	}

	return
}

func doesExistBinSearch(a []int, term int) bool {
	a = insertionSort(a)
	half := len(a) / int(2)

	for len(a) > 0 {
		fmt.Println(a)
		if term < a[half] {
			a = a[:half]
		} else if term > a[half] {
			a = a[half:]
		} else {
			return true
		}

		half = len(a) / 2
	}

	return false
}
