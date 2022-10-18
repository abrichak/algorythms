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

	sortInsertion(&src)

	fmt.Printf("%v\n", src)
}

func sortInsertion(src *[]int32) {
	result := *src
	elementsNumber := len(result)

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
