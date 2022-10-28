package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	bestCase = iota
	averageCase
	worstCase
)

const N = 100

func main() {
	src := prepareSlice(N, averageCase)
	fmt.Printf("%v\n", src)

	//sortInsertion(&src, 0, len(src)-1)
	//sortMerge(&src, 0, len(src)-1)
	sortMergeWithGoroutines(&src, 0, len(src)-1)

	fmt.Printf("%v\n", src)
}

// Insertion sort: asymptotic approximations:
// O(sqr(N))
// ~ sqr(N)
// o(N)
func sortInsertion(src *[]int32, startIndex int, endIndex int) {
	result := *src

	// Loop invariant: at the beginning of every iteration, the sub-array result[startIndex, i-1] consists of elements
	// that were initially in result[startIndex, i-1] and now are sorted
	for i := startIndex + 1; i <= endIndex; i++ {
		current := result[i]

		for j := i - 1; j >= startIndex; j-- {
			if result[j] <= current {
				break
			}

			result[j+1] = result[j]
			result[j] = current
		}
	}
}

// Merge sort: asymptotic approximations:
// ~ (N * lg N)  (for all cases)
func sortMerge(src *[]int32, startIndex int, endIndex int) {
	if startIndex < endIndex {
		centrumIndex := (startIndex + endIndex) / 2

		sortMerge(src, startIndex, centrumIndex)
		sortMerge(src, centrumIndex+1, endIndex)

		merge(src, startIndex, centrumIndex, endIndex)
	}
}

func sortMergeWithGoroutines(src *[]int32, startIndex int, endIndex int) {
	if startIndex < endIndex {
		centrumIndex := (startIndex + endIndex) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			sortMerge(src, startIndex, centrumIndex)
		}()

		go func() {
			defer wg.Done()
			sortMerge(src, centrumIndex+1, endIndex)
		}()

		wg.Wait()
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
