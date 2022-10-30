package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkSortingBestN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(N, bestCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverageN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(N, averageCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorstN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(N, worstCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverageN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(N, averageCase)
	}
}

func BenchmarkSortingBest10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*N, bestCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*N, averageCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(10*N, worstCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage10N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(10*N, averageCase)
	}
}

func BenchmarkSortingBest100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*N, bestCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*N, averageCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(100*N, worstCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage100N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(100*N, averageCase)
	}
}

func BenchmarkSortingBest1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*N, bestCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingAverage1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*N, averageCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkSortingWorst1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareSlice(1000*N, worstCase)
		//sortInsertion(&src, 0, len(src)-1)
		//sortMerge(&src, 0, len(src)-1)
		//sortMergeWithGoroutines(&src, 0, len(src)-1)
		sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)
	}
}

func BenchmarkOnlyPrepareSliceAverage1000N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareSlice(1000*N, averageCase)
	}
}

func TestSortInsertion(t *testing.T) {
	src := make([]int32, N)
	for i := 0; i < N; i++ {
		src[i] = int32(N - i - 1)
	}

	sortInsertion(&src, 0, N-1)

	for i := 0; i < N; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}

func TestSortMerge(t *testing.T) {
	src := make([]int32, N)
	for i := 0; i < N; i++ {
		src[i] = int32(N - i - 1)
	}

	sortMerge(&src, 0, len(src)-1)

	for i := 0; i < N; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}

func TestSortMergeWithGoroutines(t *testing.T) {
	src := make([]int32, N)
	for i := 0; i < N; i++ {
		src[i] = int32(N - i - 1)
	}

	sortMergeWithGoroutines(&src, 0, len(src)-1)

	for i := 0; i < N; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}

func TestSortMergeCombinedWithSortInsertion(t *testing.T) {
	src := make([]int32, N)
	for i := 0; i < N; i++ {
		src[i] = int32(N - i - 1)
	}

	sortMergeCombinedWithSortInsertion(&src, 0, len(src)-1)

	for i := 0; i < N; i++ {
		assert.Equal(t, int32(i), src[i])
	}
}
