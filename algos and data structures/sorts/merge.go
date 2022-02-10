package main

import "fmt"

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func main() {
	a := []int{3, 41, 52, 26, 38, 57, 9, 49}

	fmt.Printf("Unsorted: %v\n", a)
	fmt.Printf("Sorted: %v\n", mergeSort(a))
}

func mergeSort(a []int) []int {
	r := len(a)

	if r == 1 {
		return a
	}

	mid := len(a) / 2
	L := a[:mid]
	R := a[mid:]

	return merge(mergeSort(L), mergeSort(R))
}

func merge(l, r []int) []int {
	merged := make([]int, len(r)+len(l))

	i := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			merged[i] = l[0]
			l = l[1:]
		} else {
			merged[i] = r[0]
			r = r[1:]
		}
		i++
	}

	for j := 0; j < len(l); j++ {
		merged[i] = l[j]
		i++
	}

	for j := 0; j < len(r); j++ {
		merged[i] = r[j]
		i++
	}

	return merged
}
