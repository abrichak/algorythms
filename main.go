package main

import (
	"fmt"
	"math/rand"
)

const (
	bestCase = iota
	averageCase
	worstCase
)

const n = 100

func main() {
	src := prepareSlice(n, averageCase)
	fmt.Printf("%v\n", src)

	//sortInsertion(&src)
	sortMerge(&src, 0, len(src)-1)

	fmt.Printf("%v\n", src)
}

// Insertion sort: asymptotic approximations:
// O(sqr(n))
// ~ sqr(n)
// o(n)
func sortInsertion(src *[]int32) {
	result := *src
	elementsNumber := len(result)

	// Loop invariant: at the beginning of every iteration, the sub-array result[0, i-1] consists of elements
	// that were initially in result[0, i-1] and now are sorted
	for i := 1; i < elementsNumber; i++ {
		current := result[i]

		for j := i - 1; j >= 0; j-- {
			if result[j] <= current {
				break
			}

			result[j+1] = result[j]
			result[j] = current
		}
	}
}

// Merge sort: asymptotic approximations:
// ~ (n * lg n)  (for all cases)
func sortMerge(src *[]int32, startIndex int, endIndex int) {
	if startIndex < endIndex {
		centrumIndex := (startIndex + endIndex) / 2

		sortMerge(src, startIndex, centrumIndex)
		sortMerge(src, centrumIndex+1, endIndex)

		merge(src, startIndex, centrumIndex, endIndex)
	}
}

func prepareSlice(elementsNumber int32, caseType int) []int32 {
	var i int32

	result := make([]int32, elementsNumber)

	for i = 0; i < elementsNumber; i++ {
		switch caseType {
		case bestCase:
			result[i] = 10 * i
		case worstCase:
			result[i] = 10*elementsNumber - 10*i
		case averageCase:
			result[i] = rand.Int31n(elementsNumber)
		}
	}

	return result
}

func merge(src *[]int32, startIndex int, centrumIndex int, endIndex int) {
	result := *src

	left := make([]int32, centrumIndex-startIndex+1)
	for i := startIndex; i <= centrumIndex; i++ {
		left[i-startIndex] = result[i]
	}

	right := make([]int32, endIndex-centrumIndex)
	for i := centrumIndex + 1; i <= endIndex; i++ {
		right[i-centrumIndex-1] = result[i]
	}

	lenLeft := len(left)
	lenRight := len(right)
	i := 0
	j := 0

	for k := startIndex; k <= endIndex; k++ {
		if i >= lenLeft || (j < lenRight && left[i] > right[j]) {
			result[k] = right[j]
			j++
		} else {
			result[k] = left[i]
			i++
		}
	}
}
