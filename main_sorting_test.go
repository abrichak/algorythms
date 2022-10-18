package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSortingBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(bestCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(averageCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(worstCase)
		sortInsertion(&src)
	}
}

func BenchmarkOnlyPrepareSliceAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(averageCase)
	}
}

func TestSliceIsSorted(t *testing.T) {
	src := make([]int32, n)
	for i := 0; i < n; i++ {
		src[i] = int32(n - i - 1)
	}

	sortInsertion(&src)

	for i := 0; i < n; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}
