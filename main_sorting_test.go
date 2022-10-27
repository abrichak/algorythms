package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSortingBestN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, bestCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverageN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, averageCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorstN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(n, worstCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
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
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*n, averageCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*n, worstCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(10*n, averageCase)
	}
}

func BenchmarkSortingBest100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*n, bestCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*n, averageCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*n, worstCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(100*n, averageCase)
	}
}

func BenchmarkSortingBest1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*n, bestCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*n, averageCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*n, worstCase)
		//sortInsertion(&src)
		sortMerge(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(1000*n, averageCase)
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
