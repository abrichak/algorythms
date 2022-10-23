package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSortingBestN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, bestCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingAverageN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, averageCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingWorstN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, worstCase)
		sortInsertion(&src)
	}
}

func BenchmarkOnlyPrepareSliceAverageN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(n, averageCase)
	}
}

func BenchmarkSortingBest10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*n, bestCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*n, averageCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingWorst10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*n, worstCase)
		sortInsertion(&src)
	}
}

func BenchmarkOnlyPrepareSliceAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(10*n, averageCase)
	}
}

func TestSortInsertion(t *testing.T) {
	src := make([]int32, n)
	for i := 0; i < n; i++ {
		src[i] = int32(n - i - 1)
	}

	sortInsertion(&src)

	for i := 0; i < n; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}

func TestSortMerge(t *testing.T) {
	src := make([]int32, n)
	for i := 0; i < n; i++ {
		src[i] = int32(n - i - 1)
	}

	sortMerge(&src, 0, len(src)-1)

	for i := 0; i < n; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}
