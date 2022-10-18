package main

import (
	"fmt"
	"math/rand"
)

const n = 100

func main() {
	src := prepareArray()
	fmt.Printf("%v\n", src)

	sortInsertion(&src)

	fmt.Printf("%v\n", src)
}

func sortInsertion(src *[]int32) {
	result := *src

	for i := 1; i < n; i++ {
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

func prepareArray() []int32 {
	result := make([]int32, n)

	for i := 0; i < n; i++ {
		result[i] = rand.Int31n(n)
	}

	return result
}
