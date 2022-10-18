package main

import (
	"fmt"
	"math/rand"
)

const n = 10

func main() {
	src := prepareArray()
	fmt.Printf("%v\n", src)

	sort(&src)

	fmt.Printf("%v\n", src)
}

func sort(*[]int32) {
}

func prepareArray() []int32 {
	result := make([]int32, n)

	for i := 0; i < n; i++ {
		result[i] = rand.Int31n(n)
	}

	return result
}
